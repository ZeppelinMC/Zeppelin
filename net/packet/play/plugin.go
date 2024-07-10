package play

import (
	"github.com/dynamitemc/aether/net/io"
	"github.com/dynamitemc/aether/net/packet/configuration"
)

// clientbound
const PacketIdClientboundPluginMessage = 0x19

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

// serverbound
const PacketIdServerboundPluginMessage = 0x12

type ServerboundPluginMessage struct {
	configuration.ServerboundPluginMessage
}

func (ServerboundPluginMessage) ID() int32 {
	return 0x12
}
