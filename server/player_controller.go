package server

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/player"
	"github.com/dynamitemc/dynamite/server/world"
)

type PlayerController struct {
	player  *player.Player
	session *Session

	UUID string
}

func (p *PlayerController) JoinDimension(d *world.Dimension) error {
	if err := p.session.SendPacket(&packet.JoinGame{
		EntityID:           0, //TODO
		IsHardcore:         p.player.IsHardcore(),
		GameMode:           p.player.GameMode(),
		PreviousGameMode:   p.player.PreviousGameMode(),
		DimensionNames:     []string{d.Type()},
		DimensionName:      d.Type(),
		DimensionType:      d.Type(),
		HashedSeed:         d.Seed(),
		ViewDistance:       p.player.ViewDistance(),
		SimulationDistance: p.player.SimulationDistance(),
	}); err != nil {
		return err
	}

	return p.session.SendPacket(&packet.SetDefaultSpawnPosition{})
}

func (p *PlayerController) SendAvailableCommands(commands *packet.DeclareCommands) error {
	return p.session.SendPacket(commands)
}

func (p *PlayerController) SystemChatMessage(s string) error {
	return p.session.SendPacket(&packet.SystemChatMessage{Content: s})
}

func (p *PlayerController) Position() (x float64, y float64, z float64) {
	return p.player.X, p.player.Y, p.player.Z
}

func (p *PlayerController) Rotation() (yaw float32, pitch float32) {
	return p.player.Yaw, p.player.Pitch
}

func (p *PlayerController) OnGround() bool {
	return p.player.OnGround
}
