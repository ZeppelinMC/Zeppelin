package server

import (
	"math"
	"math/rand"
	"slices"
	"sync"

	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/player"
	"github.com/dynamitemc/dynamite/server/world"
)

type PlayerController struct {
	mu      sync.RWMutex
	player  *player.Player
	session *Session
	Server  *Server

	spawnedEntities []int32

	UUID string
}

func (p *PlayerController) Name() string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.session.conn.Info.Name
}

func (p *PlayerController) JoinDimension(d *world.Dimension) error {
	if err := p.session.SendPacket(&packet.JoinGame{
		EntityID:            p.player.EntityId(),
		IsHardcore:          p.player.IsHardcore(),
		GameMode:            p.player.GameMode(),
		PreviousGameMode:    p.player.PreviousGameMode(),
		DimensionNames:      []string{d.Type()},
		DimensionType:       d.Type(),
		DimensionName:       d.Type(),
		HashedSeed:          d.Seed(),
		MaxPlayers:          0,
		ViewDistance:        p.player.ViewDistance(),
		SimulationDistance:  p.player.SimulationDistance(),
		ReducedDebugInfo:    false,
		EnableRespawnScreen: false,
		IsDebug:             false,
		IsFlat:              false,
		DeathDimensionName:  "",
		DeathLocation:       0,
		PartialCooldown:     0,
	}); err != nil {
		return err
	}
	p.session.SendPacket(&packet.PluginMessage{
		Channel: "minecraft:brand",
		Data:    []byte("Dynamite 1.20.1"),
	})

	p.Teleport(9, 86, 8, 0, 0)

	p.session.SendPacket(&packet.SetCenterChunk{})
	p.SendSpawnChunks()

	p.SetGameMode(1)

	return p.session.SendPacket(&packet.SetDefaultSpawnPosition{})
}

func (p *PlayerController) SystemChatMessage(s string) error {
	return p.session.SendPacket(&packet.SystemChatMessage{Content: s})
}

func (p *PlayerController) ClientSettings() player.ClientInformation {
	return p.player.ClientSettings()
}

func (p *PlayerController) Position() (x, y, z float64) {
	return p.player.Position()
}

func (p *PlayerController) Rotation() (yaw, pitch float32) {
	return p.player.Rotation()
}

func (p *PlayerController) OnGround() bool {
	return p.player.OnGround()
}

func (p *PlayerController) GameMode() byte {
	return p.player.GameMode()
}

func (p *PlayerController) SetGameMode(gm byte) {
	p.player.SetGameMode(gm)
	p.session.SendPacket(&packet.GameEvent{
		Event: 3,
		Value: float32(gm),
	})
}

func (p *PlayerController) Teleport(x, y, z float64, yaw, pitch float32) {
	p.Server.teleportCounter++
	p.session.SendPacket(&packet.PlayerPositionLook{
		X:          x,
		Y:          y,
		Z:          z,
		Yaw:        yaw,
		Pitch:      pitch,
		TeleportID: p.Server.teleportCounter,
	})
}

// for now ig

func (p *PlayerController) SendCommands(graph commands.Graph) {
	for i, command := range graph.Commands {
		if command == nil {
			continue
		}
		if !p.HasPermissions(command.RequiredPermissions) {
			graph.Commands[i] = nil
		}
	}
	p.session.SendPacket(graph.Data())
}

func (p *PlayerController) Keepalive() {
	id := rand.Int63() * 100
	p.session.SendPacket(&packet.KeepAlive{PayloadID: id})
}

func (p *PlayerController) Disconnect(reason string) {
	pk := &packet.DisconnectPlay{}
	pk.Reason = reason
	p.session.SendPacket(pk)
}

func (p *PlayerController) SendSpawnChunks() {

	for x := int32(-10); x < 10; x++ {
		for z := int32(-10); z < 10; z++ {
			x, z = 0, 0
			c, err := p.Server.world.DefaultDimension().Chunk(x, z)
			if err != nil {
				panic(err)
			}

			p.session.SendPacket(c.Data())
			return
		}
	}
}

func (p *PlayerController) Chat(message string) {
	prefix, suffix := p.GetPrefixSuffix()
	msg := p.Server.Translate(p.Server.Config.Chat.Format, map[string]string{
		"player":        p.Name(),
		"player_prefix": prefix,
		"player_suffix": suffix,
		"message":       message,
	})
	p.Server.GlobalMessage(msg, p)
}

func (p *PlayerController) GetPrefixSuffix() (prefix string, suffix string) {
	group := getGroup(getPlayer(p.UUID).Group)
	return group.Prefix, group.Suffix
}

func (p *PlayerController) HandleCenterChunk(x1, z1, x2, z2 float64) {
	oldChunkX := int(math.Floor(x1 / 16))
	oldChunkZ := int(math.Floor(z1 / 16))

	newChunkX := int(math.Floor(x2 / 16))
	newChunkZ := int(math.Floor(z2 / 16))

	if newChunkX != oldChunkX || newChunkZ != oldChunkZ {
		p.session.SendPacket(&packet.SetCenterChunk{
			ChunkX: int32(newChunkX),
			ChunkZ: int32(newChunkZ),
		})
	}
}

func (p *PlayerController) IsSpawned(entityId int32) bool {
	p.mu.RLock()
	defer p.mu.RUnlock()
	for _, e := range p.spawnedEntities {
		if e == entityId {
			return true
		}
	}
	return false
}

func (p *PlayerController) SpawnPlayer(pl *PlayerController) {
	p.mu.Lock()
	defer p.mu.Unlock()
	entityId := pl.player.EntityId()
	x, y, z := pl.player.Position()
	yaw, pitch := pl.player.Rotation()

	p.session.SendPacket(&packet.SpawnPlayer{
		EntityID:   entityId,
		PlayerUUID: pl.session.Info().UUID,
		X:          x,
		Y:          y,
		Z:          z,
		Yaw:        byte(yaw),
		Pitch:      byte(pitch),
	})
	p.spawnedEntities = append(p.spawnedEntities, entityId)
}

func (p *PlayerController) DespawnPlayer(pl *PlayerController) {
	p.mu.Lock()
	defer p.mu.Unlock()
	entityId := pl.player.EntityId()

	p.session.SendPacket(&packet.DestroyEntities{
		EntityIds: []int32{entityId},
	})
	index := -1
	for i, e := range p.spawnedEntities {
		if e == entityId {
			index = i
		}
	}
	if index > -1 {
		p.spawnedEntities = slices.Delete(p.spawnedEntities, index, index+1)
	}
}
