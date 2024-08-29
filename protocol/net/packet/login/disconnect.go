package login

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io"
	"github.com/zeppelinmc/zeppelin/protocol/text"
)

// clientbound
const PacketIdDisconnect = 0x00

type Disconnect struct {
	Reason text.TextComponent
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
