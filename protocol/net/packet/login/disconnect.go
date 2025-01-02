package login

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
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

func (d *Disconnect) Encode(w encoding.Writer) error {
	return w.JSONTextComponent(d.Reason)
}

func (d *Disconnect) Decode(r encoding.Reader) error {
	return r.JSONTextComponent(&d.Reason)
}
