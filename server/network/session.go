package network

import (
	"github.com/aimjel/minecraft"
	"github.com/aimjel/minecraft/packet"
	player2 "github.com/aimjel/minecraft/player"
	"github.com/dynamitemc/dynamite/server/network/packets"
	"github.com/dynamitemc/dynamite/server/player"
	"net"
)

type Session struct {
	conn *minecraft.Conn

	state *player.Player
}

func New(c *minecraft.Conn, s *player.Player) *Session {
	return &Session{conn: c, state: s}
}

func (s *Session) HandlePackets() error {
	for {
		p, err := s.conn.ReadPacket()
		if err != nil {
			return err
		}

		switch pk := p.(type) {
		case *packet.ChatMessageServer:
			packets.ChatMessagePacket(pk.Message)
		case *packet.ChatCommandServer:
			packets.ChatCommandPacket(pk.Command)
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
