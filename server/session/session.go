package session

import (
	"aether/atomic"
	"aether/net"
	"aether/net/io"
	"aether/net/packet/configuration"
	"aether/net/packet/play"
	"aether/server/player"
	"aether/server/world"
	"bytes"
)

type Session struct {
	world  *world.World
	player *player.Player

	conn *net.Conn

	clientName string // constant
	ClientInfo atomic.AtomicValue[configuration.ClientInformation]
}

func NewSession(conn *net.Conn, entityId int32, world *world.World) *Session {
	return &Session{
		conn:   conn,
		world:  world,
		player: player.NewPlayer(entityId),
	}
}

func (session *Session) Login() error {
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

	if err := session.conn.WritePacket(&play.ClientboundPluginMessage{
		Channel: "minecraft:brand",
		Data:    io.AppendString(nil, "Aether"),
	}); err != nil {
		return err
	}

	if err := session.conn.WritePacket(&play.GameEvent{Event: play.GameEventStartWaitingChunks}); err != nil {
		return err
	}

	return nil
}

func (session *Session) sendSpawnChunks() error {
	viewDistance := int32(session.ClientInfo.Get().ViewDistance)
	var buf = new(bytes.Buffer)

	for x := 0 - viewDistance; x <= 0+viewDistance; x++ {
		for z := 0 - viewDistance; z < 0+viewDistance; z++ {
			c, err := session.world.GetChunk(x, z)
			if err != nil {
				return err
			}

			if err := session.conn.WritePacket(c.Encode(buf)); err != nil {
				return err
			}
			buf.Reset()
		}
	}

	return nil
}
