package std

import (
	"bytes"
	nnet "net"
	"slices"
	"sync"

	"github.com/dynamitemc/aether/atomic"
	"github.com/dynamitemc/aether/chat"
	"github.com/dynamitemc/aether/net"
	"github.com/dynamitemc/aether/net/io"
	"github.com/dynamitemc/aether/net/packet/configuration"
	"github.com/dynamitemc/aether/net/packet/login"
	"github.com/dynamitemc/aether/net/packet/play"
	"github.com/dynamitemc/aether/server/player"
	"github.com/dynamitemc/aether/server/session"
	"github.com/dynamitemc/aether/server/world"
	"github.com/google/uuid"
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

	conn *net.Conn

	clientName string // constant

	hasSessionData atomic.AtomicValue[bool]
	sessionData    atomic.AtomicValue[play.PlayerSession]

	lastKeepAlive int64

	spawned_ents_mu sync.Mutex
	spawnedEntities []int32
}

func NewStandardSession(conn *net.Conn, player *player.Player, world *world.World, broadcast *session.Broadcast) *StandardSession {
	return &StandardSession{
		conn:      conn,
		world:     world,
		player:    player,
		broadcast: broadcast,
	}
}

func (session *StandardSession) Teleport(x, y, z float64, yaw, pitch float32) error {
	session.player.SetPosition(x, y, z)
	session.player.SetRotation(yaw, pitch)
	return session.conn.WritePacket(&play.SynchronizePlayerPosition{X: x, Y: y, Z: z, Yaw: yaw, Pitch: pitch})
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
		SenderName: chat.TextComponent{Text: pk.Message},
	})
	return nil
}

func (session *StandardSession) PlayerInfoUpdate(pk *play.PlayerInfoUpdate) error {
	return session.conn.WritePacket(pk)
}

func (session *StandardSession) PlayerInfoRemove(uuids ...uuid.UUID) error {
	return session.conn.WritePacket(&play.PlayerInfoRemove{UUIDs: uuids})
}

func (session *StandardSession) Disconnect(reason chat.TextComponent) error {
	if session.inConfiguration() {
		return session.conn.WritePacket(&configuration.Disconnect{Reason: reason})
	} else {
		return session.conn.WritePacket(&play.Disconnect{Reason: reason})
	}
}

func (session *StandardSession) Login() error {
	go session.handlePackets()
	for _, packet := range configuration.ConstructRegistryPackets() {
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

		ViewDistance:        12,
		SimulationDistance:  12,
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
		Data:    io.AppendString(nil, "Aether"),
	}); err != nil {
		return err
	}

	if err := session.conn.WritePacket(&play.GameEvent{Event: play.GameEventStartWaitingChunks}); err != nil {
		return err
	}

	if err := session.conn.WritePacket(updateTags); err != nil {
		return err
	}

	session.Teleport(13, 65, 7, 0, 0)

	session.broadcast.AddPlayer(session)
	session.broadcast.SpawnPlayer(session)

	return nil
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

func (session *StandardSession) sendSpawnChunks() error {
	viewDistance := int32(session.Player().ClientInformation().ViewDistance)
	var buf = new(bytes.Buffer)

	if err := session.conn.WritePacket(&play.ChunkBatchStart{}); err != nil {
		return err
	}

	var chunks int32
	for x := 0 - viewDistance; x <= 0+viewDistance; x++ {
		for z := 0 - viewDistance; z < 0+viewDistance; z++ {
			c, err := session.world.GetChunk(x, z)
			if err != nil {
				continue
			}

			if err := session.conn.WritePacket(c.Encode(buf)); err != nil {
				return err
			}
			buf.Reset()
			chunks++
		}
	}

	if err := session.conn.WritePacket(&play.ChunkBatchFinished{BatchSize: chunks}); err != nil {
		return err
	}

	return nil
}
