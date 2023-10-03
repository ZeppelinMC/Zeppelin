package server

import (
	"net"

	"github.com/aimjel/minecraft"
	"github.com/aimjel/minecraft/packet"
	player2 "github.com/aimjel/minecraft/player"
	"github.com/dynamitemc/dynamite/server/network/handlers"
	"github.com/dynamitemc/dynamite/server/player"
)

type Session struct {
	conn *minecraft.Conn

	state *player.Player
}

func New(c *minecraft.Conn, s *player.Player) *Session {
	return &Session{conn: c, state: s}
}

func (s *Session) HandlePackets(controller *PlayerController) error {
	for {
		p, err := s.conn.ReadPacket()
		if err != nil {
			return err
		}

		switch pk := p.(type) {
		case *packet.ChatMessageServer:
			handlers.ChatMessagePacket(controller, pk.Message)
		case *packet.ChatCommandServer:
			handlers.ChatCommandPacket(controller, controller.Server.CommandGraph, pk.Command)
		case *packet.ClientSettings:
			handlers.ClientSettings(s.state, pk)
		}
		switch p.ID() {
		case 14, 15, 16, 17:
			{
				handlers.PlayerMovement(controller, s.state, p)
			}
		}
	}
}

func (s *Session) SendPacket(p packet.Packet) error {
	return s.conn.SendPacket(p)
}

func (s *Session) Info() *player2.Info {
	return s.conn.Info
}

func (s *Session) RemoteAddr() net.Addr {
	return s.conn.RemoteAddr()
}
