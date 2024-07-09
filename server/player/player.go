package player

import (
	"aether/atomic"
	"aether/net"
	"aether/net/io"
	"aether/net/packet/configuration"
	"aether/net/packet/play"
	"aether/net/registry"
	"aether/server/world"
)

type Player struct {
	world *world.World

	conn     *net.Conn
	entityId int32 // constant

	clientName string // constant
	clientInfo atomic.AtomicValue[*configuration.ClientInformation]
}

func NewPlayer(conn *net.Conn, entityId int32, world *world.World) *Player {
	return &Player{
		conn:     conn,
		entityId: entityId,
		world:    world,
	}
}

func (p *Player) Login() error {
	go p.handlePackets()
	for _, packet := range registry.RegistryMap.Packets() {
		if err := p.conn.WritePacket(packet); err != nil {
			return err
		}
	}
	if err := p.conn.WritePacket(configuration.FinishConfiguration{}); err != nil {
		return err
	}

	if err := p.conn.WritePacket(&play.Login{
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

	if err := p.conn.WritePacket(&play.ClientboundPluginMessage{
		Channel: "minecraft:brand",
		Data:    io.AppendString(nil, "Aether"),
	}); err != nil {
		return err
	}

	if err := p.conn.WritePacket(&play.GameEvent{Event: play.GameEventStartWaitingChunks}); err != nil {
		return err
	}

	for x, z := int32(-6), int32(-6); x < 6 && z < 6; x, z = x+1, z+1 {
		c, _ := p.world.GetChunk(x, z)

		p.conn.WritePacket(c.Encode())
	}

	return nil
}
