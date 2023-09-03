package network

import (
	"github.com/aimjel/minecraft"
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server/network/packets"
)

type Session struct {
	Conn *minecraft.Conn
}

type server interface {
	PlayerlistRemove(players ...[16]byte)
	GlobalBroadcast(packet.Packet)
}

func NewSession(c *minecraft.Conn, srv server, logger logger.Logger) *Session {
	go func() {
		for {
			pk, err := c.ReadPacket()
			if err != nil {
				packets.Disconnect(c, srv, logger)
				break
			}
			switch pk := pk.(type) {
			case *packet.ChatMessageServer:
				{
					packets.ChatMessagePacket(pk.Message, srv)
				}
			}
		}
	}()
	return &Session{Conn: c}
}
