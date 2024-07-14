package std

import (
	"bytes"
	nnet "net"

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

// StandardSession is a session that uses *net.Conn
type StandardSession struct {
	world     *world.World
	player    *player.Player
	broadcast *session.Broadcast

	conn *net.Conn

	clientName string // constant

	hasSessionData atomic.AtomicValue[bool]
	sessionData    atomic.AtomicValue[play.PlayerSession]
}

func NewStandardSession(conn *net.Conn, entityId int32, world *world.World, broadcast *session.Broadcast) *StandardSession {
	return &StandardSession{
		conn:      conn,
		world:     world,
		player:    player.NewPlayer(entityId),
		broadcast: broadcast,
	}
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
		panic("didnt implement play disconnect")
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

	session.broadcast.AddPlayer(session)

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
