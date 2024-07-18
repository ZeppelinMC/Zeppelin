package play

import (
	"github.com/dynamitemc/aether/net/io"
	"github.com/dynamitemc/aether/text"
)

// clientbound
const PacketIdDisconnect = 0x1D

type Disconnect struct {
	Reason text.TextComponent
}

func (Disconnect) ID() int32 {
	return 0x1D
}

func (d *Disconnect) Encode(w io.Writer) error {
	return w.TextComponent(d.Reason)
}

func (d *Disconnect) Decode(r io.Reader) error {
	return r.TextComponent(&d.Reason)
}
