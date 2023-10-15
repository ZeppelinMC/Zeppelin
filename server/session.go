package server

import (
	"errors"
	"io"
	"net"

	"github.com/aimjel/minecraft"
	"github.com/aimjel/minecraft/packet"
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
		if errors.Is(err, io.EOF) {
			return err
		}

		switch pk := p.(type) {
		case *packet.PlayerCommandServer:
			handlers.PlayerCommand(controller, pk.ActionID)
		case *packet.ChatMessageServer:
			handlers.ChatMessagePacket(controller, pk.Message)
		case *packet.ChatCommandServer:
			handlers.ChatCommandPacket(controller, controller.Server.commandGraph, pk.Command)
		case *packet.ClientSettings:
			handlers.ClientSettings(controller, s.state, pk)
		case *packet.PlayerPosition, *packet.PlayerPositionRotation, *packet.PlayerRotation:
			handlers.PlayerMovement(controller, s.state, p)
		case *packet.PlayerActionServer:
			handlers.PlayerAction(controller, pk)
		case *packet.InteractServer:
			handlers.Interact(controller, pk)
		case *packet.SwingArmServer:
			handlers.SwingArm(controller, pk.Hand)
		case *packet.CommandSuggestionsRequest:
			handlers.CommandSuggestionsRequest(pk.TransactionId, pk.Text, controller.Server.commandGraph, controller)
		case *packet.ClientCommandServer:
			handlers.ClientCommand(controller, s.state, pk.ActionID)
		case *packet.PlayerAbilitiesServer:
			handlers.PlayerAbilities(s.state, pk.Flags)
		}
	}
}

func (s *Session) SendPacket(p packet.Packet) error {
	return s.conn.SendPacket(p)
}

func (s *Session) RemoteAddr() net.Addr {
	return s.conn.RemoteAddr()
}
