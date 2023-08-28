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
	go func() {
		for {
			pk, err := c.ReadPacket()
			if err != nil {
				break
			}
			switch pk := pk.(type) {
			case *packet.ChatMessageServer:
				{
					packets.ChatMessagePacket(pk.Message)
				}
			}
		}
	}()
	return &Session{Conn: c}
}
