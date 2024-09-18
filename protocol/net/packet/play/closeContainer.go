package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// serverbound
const PacketIdCloseContainer = 0x0F

type CloseContainer struct {
	WindowId byte
}

func (CloseContainer) ID() int32 {
	return PacketIdCloseContainer
}

func (c *CloseContainer) Encode(w encoding.Writer) error {
	return w.Ubyte(c.WindowId)
}

func (c *CloseContainer) Decode(r encoding.Reader) error {
	return r.Ubyte(&c.WindowId)
}
