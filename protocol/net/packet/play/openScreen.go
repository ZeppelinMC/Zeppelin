package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io"
	"github.com/zeppelinmc/zeppelin/protocol/text"
)

// clientbound
const PacketIdOpenScreen = 0x33

type OpenScreen struct {
	WindowId    int32
	WindowType  int32
	WindowTitle text.TextComponent
}

func (OpenScreen) ID() int32 {
	return PacketIdOpenScreen
}

func (o *OpenScreen) Encode(w io.Writer) error {
	if err := w.VarInt(o.WindowId); err != nil {
		return err
	}
	if err := w.VarInt(o.WindowType); err != nil {
		return err
	}
	return w.TextComponent(o.WindowTitle)
}

func (o *OpenScreen) Decode(r io.Reader) error {
	if _, err := r.VarInt(&o.WindowId); err != nil {
		return err
	}
	if _, err := r.VarInt(&o.WindowType); err != nil {
		return err
	}
	return r.TextComponent(&o.WindowTitle)
}
