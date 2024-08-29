package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io"
	"github.com/zeppelinmc/zeppelin/protocol/text"
)

// clientbound
const PacketIdServerData = 0x4B

type ServerData struct {
	MOTD text.TextComponent
	Icon []byte
}

func (ServerData) ID() int32 {
	return PacketIdServerData
}

func (d *ServerData) Encode(w io.Writer) error {
	if err := w.TextComponent(d.MOTD); err != nil {
		return err
	}
	if err := w.Bool(d.Icon != nil); err != nil {
		return err
	}
	if d.Icon != nil {
		return w.ByteArray(d.Icon)
	}
	return nil
}

func (d *ServerData) Decode(r io.Reader) error {
	if err := r.TextComponent(&d.MOTD); err != nil {
		return err
	}
	var hasIcon bool
	if err := r.Bool(&hasIcon); err != nil {
		return err
	}
	if hasIcon {
		return r.ByteArray(&d.Icon)
	}
	return nil
}
