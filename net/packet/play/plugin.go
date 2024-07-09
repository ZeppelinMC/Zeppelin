package play

import (
	"aether/net/io"
	"aether/net/packet/configuration"
)

type ClientboundPluginMessage configuration.ClientboundPluginMessage

func (ClientboundPluginMessage) ID() int32 {
	return 0x19
}

func (c *ClientboundPluginMessage) Encode(w io.Writer) error {
	if err := w.Identifier(c.Channel); err != nil {
		return err
	}
	return w.FixedByteArray(c.Data)
}

func (c *ClientboundPluginMessage) Decode(r io.Reader) error {
	if err := r.Identifier(&c.Channel); err != nil {
		return err
	}
	return r.ReadAll(&c.Data)
}

type ServerboundPluginMessage struct {
	configuration.ServerboundPluginMessage
}

func (ServerboundPluginMessage) ID() int32 {
	return 0x12
}
