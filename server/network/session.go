package network

import (
	"github.com/aimjel/minecraft"
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/network/packets"
)

type Session struct {
	Conn *minecraft.Conn
}

func NewSession(c *minecraft.Conn) *Session {
	return &Session{Conn: c}
}

func (s *Session) HandlePackets() error {
	for {
		p, err := s.Conn.ReadPacket()
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
