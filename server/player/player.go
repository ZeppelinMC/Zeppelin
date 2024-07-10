package player

import (
	"aether/atomic"
	"aether/net"
	"aether/net/io"
	"aether/net/packet/configuration"
	"aether/net/packet/play"
	"aether/server/world"
)

type Player struct {
	world *world.World

	conn     *net.Conn
	entityId int32 // constant

	clientName string // constant
	clientInfo atomic.AtomicValue[configuration.ClientInformation]
}

func NewPlayer(conn *net.Conn, entityId int32, world *world.World) *Player {
	return &Player{
		conn:     conn,
		entityId: entityId,
		world:    world,
	}
}

func (player *Player) Login() error {
	go player.handlePackets()
	for _, packet := range configuration.ConstructRegistryPackets() {
		if err := player.conn.WritePacket(packet); err != nil {
			return err
		}
	}
	if err := player.conn.WritePacket(configuration.FinishConfiguration{}); err != nil {
		return err
	}

	if err := player.conn.WritePacket(&play.Login{
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

	if err := player.conn.WritePacket(&play.ClientboundPluginMessage{
		Channel: "minecraft:brand",
		Data:    io.AppendString(nil, "Aether"),
	}); err != nil {
		return err
	}

	if err := player.conn.WritePacket(&play.GameEvent{Event: play.GameEventStartWaitingChunks}); err != nil {
		return err
	}

	return nil
}

func (player *Player) sendSpawnChunks() error {
	viewDistance := int32(player.clientInfo.Get().ViewDistance)

	for x := 0 - viewDistance; x <= 0+viewDistance; x++ {
		for z := 0 - viewDistance; z < 0+viewDistance; z++ {
			c, err := player.world.GetChunk(x, z)
			if err != nil {
				return err
			}

			if err := player.conn.WritePacket(c.Encode()); err != nil {
				return err
			}
		}
	}

	return nil
}
