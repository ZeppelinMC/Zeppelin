package configuration

import (
	"aether/chat"
	"aether/net/io"
)

type Disconnect struct {
	Reason chat.TextComponent
}

func (Disconnect) ID() int32 {
	return 0x02
}

func (d *Disconnect) Encode(w io.Writer) error {
	return nil //TODO
}

func (d *Disconnect) Decode(r io.Reader) error {
	return nil //TODO
}
