package session

import (
	"bytes"

	"github.com/dynamitemc/aether/atomic"
	"github.com/dynamitemc/aether/net"
	"github.com/dynamitemc/aether/net/io"
	"github.com/dynamitemc/aether/net/packet/configuration"
	"github.com/dynamitemc/aether/net/packet/play"
	"github.com/dynamitemc/aether/server/player"
	"github.com/dynamitemc/aether/server/world"
)

type Session struct {
	world  *world.World
	Player *player.Player

	Conn *net.Conn

	clientName string // constant
	ClientInfo atomic.AtomicValue[configuration.ClientInformation]
}

func NewSession(conn *net.Conn, entityId int32, world *world.World) *Session {
	return &Session{
		Conn:   conn,
		world:  world,
		Player: player.NewPlayer(entityId),
	}
}

func (session *Session) Login() error {
	go session.handlePackets()
	for _, packet := range configuration.ConstructRegistryPackets() {
		if err := session.Conn.WritePacket(packet); err != nil {
			return err
		}
	}
	if err := session.Conn.WritePacket(configuration.FinishConfiguration{}); err != nil {
		return err
	}

	if err := session.Conn.WritePacket(&play.Login{
		EntityID:   1,
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

	if err := session.Conn.WritePacket(&play.ClientboundPluginMessage{
		Channel: "minecraft:brand",
		Data:    io.AppendString(nil, "Aether"),
	}); err != nil {
		return err
	}

	if err := session.Conn.WritePacket(&play.GameEvent{Event: play.GameEventStartWaitingChunks}); err != nil {
		return err
	}

	return nil
}

func (session *Session) sendSpawnChunks() error {
	viewDistance := int32(session.ClientInfo.Get().ViewDistance)
	var buf = new(bytes.Buffer)

	if err := session.Conn.WritePacket(&play.ChunkBatchStart{}); err != nil {
		return err
	}

	var chunks int32
	for x := 0 - viewDistance; x <= 0+viewDistance; x++ {
		for z := 0 - viewDistance; z < 0+viewDistance; z++ {
			c, err := session.world.GetChunk(x, z)
			if err != nil {
				continue
			}

			if err := session.Conn.WritePacket(c.Encode(buf)); err != nil {
				return err
			}
			buf.Reset()
			chunks++
		}
	}

	if err := session.Conn.WritePacket(&play.ChunkBatchFinished{BatchSize: chunks}); err != nil {
		return err
	}

	return nil
}
