package configuration

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
	"github.com/zeppelinmc/zeppelin/protocol/text"
)

// clientbound
const PacketIdDisconnect = 0x02

type Disconnect struct {
	Reason text.TextComponent
}

func (Disconnect) ID() int32 {
	return 0x02
}

func (d *Disconnect) Encode(w encoding.Writer) error {
	return w.TextComponent(d.Reason)
}

func (d *Disconnect) Decode(r encoding.Reader) error {
	return r.TextComponent(&d.Reason)
}
