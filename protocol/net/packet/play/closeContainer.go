package play

import "github.com/zeppelinmc/zeppelin/protocol/net/io"

// serverbound
const PacketIdCloseContainer = 0x0F

type CloseContainer struct {
	WindowId byte
}

func (CloseContainer) ID() int32 {
	return PacketIdCloseContainer
}

func (c *CloseContainer) Encode(w io.Writer) error {
	return w.Ubyte(c.WindowId)
}

func (c *CloseContainer) Decode(r io.Reader) error {
	return r.Ubyte(&c.WindowId)
}
