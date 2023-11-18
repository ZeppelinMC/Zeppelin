package session

import (
	"net"

	"github.com/aimjel/minecraft/packet"
	"github.com/aimjel/minecraft/protocol/types"
)

type Session interface {
	SendPacket(pk packet.Packet) error
	ReadPacket() (packet.Packet, error)
	Close(err error)
	Name() string
	UUID() [16]byte
	Properties() []types.Property
	RemoteAddr() net.Addr
}
