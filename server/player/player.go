package player

import (
	"github.com/zeppelinmc/zeppelin/properties"
	"github.com/zeppelinmc/zeppelin/protocol/net"
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/configuration"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/play"
	"github.com/zeppelinmc/zeppelin/protocol/net/tags"
	"github.com/zeppelinmc/zeppelin/protocol/text"
	"github.com/zeppelinmc/zeppelin/server/command"
	"github.com/zeppelinmc/zeppelin/server/entity"
	"github.com/zeppelinmc/zeppelin/server/player/state"
	"github.com/zeppelinmc/zeppelin/server/world/dimension"
	"github.com/zeppelinmc/zeppelin/server/world/level"
	"slices"
	"sync/atomic"
	"unsafe"
)

type Player struct {
	*net.Conn
	*state.PlayerEntity

	playerList *PlayerList
	// the entity id for this player
	entityId int32

	Unlisted bool

	ClientInformation atomic.Pointer[configuration.ClientInformation]

	dimensionManager *dimension.DimensionManager
	worldLevel       *level.Level

	serverProperties *properties.ServerProperties
	commandManager   *command.Manager

	// the time in milliseconds that the keep alive packet was sent to the server from the client
	sbLastKeepalive atomic.Int64
	// the time in milliseconds that the keep alive packet was sent to the client from the server
	cbLastKeepAlive atomic.Int64

	registryIndexes map[string][]string

	// the packet currently awaited
	packetAwaited     atomic.Int32
	packetAwaitedChan atomic.Pointer[chan packet.Decodeable]
}

func int32toAtomic(i int32) atomic.Int32 {
	// int32 and atomic.Int32 are the same structure and size
	return *(*atomic.Int32)(unsafe.Pointer(&i))
}

func (list *PlayerList) New(conn *net.Conn, en *state.PlayerEntity, dimensionManager *dimension.DimensionManager, worldLevel *level.Level, serverProperties *properties.ServerProperties, commandManager *command.Manager) *Player {
	p := &Player{
		entityId:     entity.NewEntityId(),
		PlayerEntity: en,

		dimensionManager: dimensionManager,
		worldLevel:       worldLevel,

		serverProperties: serverProperties,
		commandManager:   commandManager,

		registryIndexes: make(map[string][]string),

		packetAwaited: int32toAtomic(-1),

		playerList: list,

		Conn: conn,
	}
	p.ClientInformation.Store(&configuration.ClientInformation{})

	return p
}

/*
ViewDistance returns the view distance of the client, in chunks, or the server's render distance if the client's view distance is bigger or not set
*/
func (p *Player) ViewDistance() int32 {
	plVd := int32(p.ClientInformation.Load().ViewDistance)
	if plVd == 0 || plVd > p.serverProperties.ViewDistance {
		return p.serverProperties.ViewDistance
	}

	return plVd
}

// writeOrKill tries to write the packet to the player and kills the connection on failure
func (p *Player) writeOrKill(pk packet.Encodeable) {
	if err := p.WritePacket(pk); err != nil {
		p.killConnection(false, "lost connection")
	}
}

// finishConfiguration finishes the configuration phase by sending the server brand, registries and tags
func (p *Player) finishConfiguration() error {
	// begin packet reading in the background
	go p.listenPackets()

	// send the server brand (shows in F3)
	if err := p.WritePacket(&configuration.ClientboundPluginMessage{
		Channel: "minecraft:brand",
		Data:    encoding.AppendString(nil, "Zeppelin"),
	}); err != nil {
		return err
	}

	// send the registry packets
	configuration.RegistryPacketsMutex.Lock()
	for _, pk := range configuration.RegistryPackets {
		if err := p.WritePacket(pk); err != nil {
			return err
		}
		p.registryIndexes[pk.RegistryId] = slices.Clone(pk.Indexes)
	}
	configuration.RegistryPacketsMutex.Unlock()

	// finish the configuration phase
	if err := p.WritePacket(configuration.FinishConfiguration{}); err != nil {
		return err
	}

	// wait for the client to acknowledge the finish configuration packet
	if _, err := p.awaitPacket(configuration.PacketIdAcknowledgeFinishConfiguration); err != nil {
		return err
	}

	// switch to play state
	p.SetState(net.PlayState)

	// send tags (blocks, fluids, etc)
	return p.WritePacket(tags.Tags)
}

// startGame finishes the initialization for the player and logs it into the world
func (p *Player) startGame() error {
	if err := p.sendCommandGraph(); err != nil {
		return err
	}

	if err := p.WritePacket(&play.Login{
		EntityID:           p.entityId,
		DimensionName:      p.DimensionName(),
		GameMode:           byte(p.GameMode()),
		DimensionType:      int32(slices.Index(p.registryIndexes["minecraft:dimension_type"], p.DimensionName())),
		ViewDistance:       p.serverProperties.ViewDistance,
		SimulationDistance: p.serverProperties.SimulationDistance,
		HashedSeed:         p.worldLevel.Data.WorldGenSettings.Seed.Hash(),
		EnforcesSecureChat: p.serverProperties.EnforceSecureProfile,
	}); err != nil {
		return err
	}

	// todo add all other stuff
	if err := p.WritePacket(&play.SetDefaultSpawnPosition{
		X:     p.worldLevel.Data.SpawnX,
		Y:     p.worldLevel.Data.SpawnY,
		Z:     p.worldLevel.Data.SpawnZ,
		Angle: p.worldLevel.Data.SpawnAngle,
	}); err != nil {
		return err
	}

	x, y, z := p.Position()
	yaw, pitch := p.Rotation()

	if err := p.SynchronizePosition(x, y, z, yaw, pitch); err != nil {
		return err
	}

	if err := p.sendSpawnChunks(); err != nil {
		return err
	}

	if err := p.SynchronizePosition(x, y, z, yaw, pitch); err != nil {
		return err
	}

	if err := p.WritePacket(&play.GameEvent{Event: play.GameEventStartWaitingChunks}); err != nil {
		return err
	}
	p.playerList.AddPlayer(p)

	return nil
}

// SynchronizePosition teleports the player to the specified coordinates
func (p *Player) SynchronizePosition(x, y, z float64, yaw, pitch float32) error {
	p.SetPosition(x, y, z)
	p.SetRotation(yaw, pitch)

	if err := p.WritePacket(&play.SynchronizePlayerPosition{X: x, Y: y, Z: z, Yaw: yaw, Pitch: pitch}); err != nil {
		return err
	}
	_, err := p.awaitPacket(play.PacketIdConfirmTeleportation)

	return err
}

// Login logs the player into the game
func (p *Player) Login() error {
	if err := p.finishConfiguration(); err != nil {
		return err
	}
	return p.startGame()
}

// SystemMessage sends a system message (unsigned) to the player
func (p *Player) SystemMessage(msg text.TextComponent) error {
	return p.WritePacket(&play.SystemChatMessage{Content: msg})
}

// sendCommandGraph sends the command graph to the player
func (p *Player) sendCommandGraph() error {
	return p.WritePacket(p.commandManager.Encode())
}
