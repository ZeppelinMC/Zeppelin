package std

import (
	"bytes"
	"math"
	nnet "net"
	"slices"
	"sync"

	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/atomic"
	"github.com/zeppelinmc/zeppelin/net"
	"github.com/zeppelinmc/zeppelin/net/io"
	"github.com/zeppelinmc/zeppelin/net/metadata"
	"github.com/zeppelinmc/zeppelin/net/packet"
	"github.com/zeppelinmc/zeppelin/net/packet/configuration"
	"github.com/zeppelinmc/zeppelin/net/packet/login"
	"github.com/zeppelinmc/zeppelin/net/packet/play"
	"github.com/zeppelinmc/zeppelin/server/config"
	"github.com/zeppelinmc/zeppelin/server/player"
	"github.com/zeppelinmc/zeppelin/server/session"
	"github.com/zeppelinmc/zeppelin/server/world"
	"github.com/zeppelinmc/zeppelin/server/world/region"
	"github.com/zeppelinmc/zeppelin/text"
	"github.com/zeppelinmc/zeppelin/util"
)

var _ session.Session = (*StandardSession)(nil)

var updateTags = &play.UpdateTags{
	Tags: map[string]map[string][]int32{
		"minecraft:fluid": {
			"minecraft:water": {1, 2},
			"minecraft:lava":  {3, 4},
		},
	},
}

// StandardSession is a session that uses *net.Conn
type StandardSession struct {
	world     *world.World
	player    *player.Player
	broadcast *session.Broadcast
	config    config.ServerConfig

	conn *net.Conn

	clientName string // constant

	hasSessionData atomic.AtomicValue[bool]
	sessionData    atomic.AtomicValue[play.PlayerSession]

	spawned_ents_mu sync.Mutex
	spawnedEntities []int32

	Spawned atomic.AtomicValue[bool]
}

func NewStandardSession(conn *net.Conn, player *player.Player, world *world.World, broadcast *session.Broadcast, config config.ServerConfig) *StandardSession {
	return &StandardSession{
		conn:      conn,
		world:     world,
		player:    player,
		broadcast: broadcast,
		config:    config,
	}
}

func (session *StandardSession) WritePacket(pk packet.Packet) error {
	return session.conn.WritePacket(pk)
}

func (session *StandardSession) ReadPacket() (packet.Packet, error) {
	return session.conn.ReadPacket()
}

func (session *StandardSession) SynchronizePosition(x, y, z float64, yaw, pitch float32) error {
	session.player.SetPosition(x, y, z)
	session.player.SetRotation(yaw, pitch)
	if err := session.conn.WritePacket(&play.SynchronizePlayerPosition{X: x, Y: y, Z: z, Yaw: yaw, Pitch: pitch}); err != nil {
		return err
	}
	return session.conn.WritePacket(&play.SetCenterChunk{ChunkX: int32(x * 16), ChunkZ: int32(z * 16)})
}

func (session *StandardSession) UpdateEntityPosition(pk *play.UpdateEntityPosition) error {
	return session.conn.WritePacket(pk)
}

// additionally sends head rotation
func (session *StandardSession) UpdateEntityPositionRotation(pk *play.UpdateEntityPositionAndRotation) error {
	if err := session.conn.WritePacket(pk); err != nil {
		return err
	}
	return session.conn.WritePacket(&play.SetHeadRotation{EntityId: pk.EntityId, HeadYaw: pk.Yaw})
}

// additionally sends head rotation
func (session *StandardSession) UpdateEntityRotation(pk *play.UpdateEntityRotation) error {
	if err := session.conn.WritePacket(pk); err != nil {
		return err
	}
	return session.conn.WritePacket(&play.SetHeadRotation{EntityId: pk.EntityId, HeadYaw: pk.Yaw})
}

func (session *StandardSession) EntityAnimation(entityId int32, animation byte) error {
	return session.conn.WritePacket(&play.EntityAnimation{EntityId: entityId, Animation: animation})
}

func (session *StandardSession) EntityMetadata(entityId int32, md metadata.Metadata) error {
	return session.conn.WritePacket(&play.SetEntityMetadata{EntityId: entityId, Metadata: md})
}

func (session *StandardSession) Conn() *net.Conn {
	return session.conn
}

func (session *StandardSession) Broadcast() *session.Broadcast {
	return session.broadcast
}

func (session *StandardSession) Player() *player.Player {
	return session.player
}

func (session *StandardSession) Addr() nnet.Addr {
	return session.conn.RemoteAddr()
}

func (session *StandardSession) ClientName() string {
	return session.clientName
}

func (session *StandardSession) Username() string {
	return session.conn.Username()
}

func (session *StandardSession) UUID() uuid.UUID {
	return session.conn.UUID()
}

func (session *StandardSession) Properties() []login.Property {
	return session.conn.Properties()
}

func (session *StandardSession) SessionData() (play.PlayerSession, bool) {
	return session.sessionData.Get(), session.hasSessionData.Get()
}

func (session *StandardSession) PlayerChatMessage(pk play.ChatMessage, sender session.Session, chatType int32) error {
	session.conn.WritePacket(&play.PlayerChatMessage{
		Sender:              sender.UUID(),
		Index:               0,
		HasMessageSignature: pk.HasSignature,
		MessageSignature:    pk.Signature,
		Message:             pk.Message,
		Timestamp:           pk.Timestamp,
		Salt:                pk.Salt,

		ChatType:   chatType,
		SenderName: text.TextComponent{Text: session.Username()},
	})
	return nil
}

func (session *StandardSession) PlayerInfoUpdate(pk *play.PlayerInfoUpdate) error {
	return session.conn.WritePacket(pk)
}

func (session *StandardSession) PlayerInfoRemove(uuids ...uuid.UUID) error {
	return session.conn.WritePacket(&play.PlayerInfoRemove{UUIDs: uuids})
}

func (session *StandardSession) Disconnect(reason text.TextComponent) error {
	if session.inConfiguration() {
		return session.conn.WritePacket(&configuration.Disconnect{Reason: reason})
	} else {
		return session.conn.WritePacket(&play.Disconnect{Reason: reason})
	}
}

func (session *StandardSession) Login() error {
	go session.handlePackets()
	for _, packet := range configuration.RegistryPackets {
		if err := session.conn.WritePacket(packet); err != nil {
			return err
		}
	}
	if err := session.conn.WritePacket(configuration.FinishConfiguration{}); err != nil {
		return err
	}

	if err := session.conn.WritePacket(&play.Login{
		EntityID:   session.player.EntityId(),
		Dimensions: []string{"minecraft:overworld"},

		ViewDistance:        session.config.RenderDistance,
		SimulationDistance:  session.config.SimulationDistance,
		EnableRespawnScreen: true,
		DimensionType:       0,
		DimensionName:       "minecraft:overworld",
		GameMode:            1,

		EnforcesSecureChat: true,
	}); err != nil {
		return err
	}

	if err := session.conn.WritePacket(&play.ClientboundPluginMessage{
		Channel: "minecraft:brand",
		Data:    io.AppendString(nil, session.config.Brand),
	}); err != nil {
		return err
	}

	if err := session.conn.WritePacket(updateTags); err != nil {
		return err
	}

	session.broadcast.AddPlayer(session)
	return nil
}

func (session *StandardSession) SystemMessage(msg text.TextComponent) error {
	return session.conn.WritePacket(&play.SystemChatMessage{Content: msg})
}

func (session *StandardSession) IsSpawned(entityId int32) bool {
	session.spawned_ents_mu.Lock()
	defer session.spawned_ents_mu.Unlock()
	for _, e := range session.spawnedEntities {
		if e == entityId {
			return true
		}
	}
	return false
}

func (session *StandardSession) Dimension() *region.Dimension {
	return session.world.Dimension(session.player.Dimension())
}

/*
The view distance of the client, in chunks

Returns the server's render distance if the client's view distance is bigger or not set
*/
func (session *StandardSession) ViewDistance() int32 {
	plVd := int32(session.player.ClientInformation().ViewDistance)
	if plVd == 0 || plVd > session.config.RenderDistance {
		return session.config.RenderDistance
	}

	return plVd
}

func (session *StandardSession) DespawnEntities(entityIds ...int32) error {
	session.spawned_ents_mu.Lock()
	defer session.spawned_ents_mu.Unlock()
	session.spawnedEntities = slices.DeleteFunc(session.spawnedEntities, func(entityId int32) bool {
		for _, e := range entityIds {
			if e == entityId {
				return true
			}
		}
		return false
	})
	return session.conn.WritePacket(&play.RemoveEntities{EntityIDs: entityIds})
}

func (session *StandardSession) SpawnEntity(pk *play.SpawnEntity) error {
	if err := session.conn.WritePacket(pk); err != nil {
		return err
	}
	session.spawned_ents_mu.Lock()
	defer session.spawned_ents_mu.Unlock()
	session.spawnedEntities = append(session.spawnedEntities, pk.EntityId)

	return nil
}

func (session *StandardSession) SpawnPlayer(ses session.Session) error {
	x, y, z := ses.Player().Position()
	yaw, pitch := ses.Player().Rotation()
	pk := &play.SpawnEntity{
		EntityId:   ses.Player().EntityId(),
		EntityUUID: ses.UUID(),
		Type:       128,
		X:          x,
		Y:          y,
		Z:          z,
		Yaw:        util.DegreesToAngle(yaw),
		Pitch:      util.DegreesToAngle(pitch),
		HeadYaw:    util.DegreesToAngle(yaw),
	}
	return session.SpawnEntity(pk)
}

func (session *StandardSession) sendSpawnChunks() error {
	if err := session.conn.WritePacket(&play.GameEvent{Event: play.GameEventStartWaitingChunks}); err != nil {
		return err
	}

	viewDistance := session.ViewDistance()
	var buf = new(bytes.Buffer)

	x, _, z := session.player.Position()
	chunkX, chunkZ := int32(math.Floor(x/16)), int32(math.Floor(z/16))

	for x := chunkX - viewDistance; x <= chunkX+viewDistance; x++ {
		for z := chunkZ - viewDistance; z < chunkZ+viewDistance; z++ {
			c, err := session.world.Overworld.GetChunk(x, z)
			if err != nil {
				continue
			}

			if err := session.conn.WritePacket(c.Encode(buf)); err != nil {
				return err
			}
			buf.Reset()
		}
	}

	return nil
}

func (session *StandardSession) spawn() error {
	if err := session.WritePacket(&play.SetDefaultSpawnPosition{
		X:     session.world.Data.SpawnX,
		Y:     session.world.Data.SpawnY,
		Z:     session.world.Data.SpawnZ,
		Angle: session.world.Data.SpawnAngle,
	}); err != nil {
		return err
	}

	x, y, z := session.player.Position()
	yaw, pitch := session.player.Rotation()
	if err := session.SynchronizePosition(x, y, z, yaw, pitch); err != nil {
		return err
	}

	if err := session.sendSpawnChunks(); err != nil {
		return err
	}

	return nil
}
