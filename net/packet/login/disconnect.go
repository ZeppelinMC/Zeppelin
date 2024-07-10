package login

import (
	"aether/chat"
	"aether/net/io"
)

// clientbound
const PacketIdDisconnect = 0x00

type Disconnect struct {
	Reason chat.TextComponent
}

func (Disconnect) ID() int32 {
	return 0x00
}

func (d *Disconnect) Encode(w io.Writer) error {
	return w.JSONTextComponent(d.Reason)
}

func (d *Disconnect) Decode(r io.Reader) error {
	return r.JSONTextComponent(&d.Reason)
}
