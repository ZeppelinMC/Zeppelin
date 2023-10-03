package server

import (
	"math/rand"

	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/player"
	"github.com/dynamitemc/dynamite/server/world"
)

type PlayerController struct {
	player  *player.Player
	session *Session
	Server  *Server

	UUID string
}

func (p *PlayerController) Name() string {
	return p.session.conn.Info.Name
}

func (p *PlayerController) JoinDimension(d *world.Dimension) error {
	if err := p.session.SendPacket(&packet.JoinGame{
		EntityID:           p.player.EntityID,
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
	p.session.SendPacket(&packet.PluginMessage{
		Channel: "minecraft:brand",
		Data:    []byte("Dynamite 1.20.1"),
	})

	p.Spawn()
	return p.session.SendPacket(&packet.SetDefaultSpawnPosition{})
}

func (p *PlayerController) SystemChatMessage(s string) error {
	return p.session.SendPacket(&packet.SystemChatMessage{Content: s})
}

func (p *PlayerController) ClientSettings() player.ClientInformation {
	return p.player.ClientSettings
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

func (p *PlayerController) GameMode() byte {
	return p.player.GameMode()
}

func (p *PlayerController) SetGameMode(gm byte) {
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

func (p *PlayerController) SendCommands(graph commands.Graph) {
	for i, command := range graph.Commands {
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
