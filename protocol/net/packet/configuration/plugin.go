package configuration

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// two-sided
const PacketIdPluginMessage = 0x01

type ClientboundPluginMessage struct {
	Channel string
	Data    []byte
}

func (ClientboundPluginMessage) ID() int32 {
	return 0x01
}

func (c *ClientboundPluginMessage) Encode(w encoding.Writer) error {
	if err := w.Identifier(c.Channel); err != nil {
		return err
	}
	return w.FixedByteArray(c.Data)
}

func (c *ClientboundPluginMessage) Decode(r encoding.Reader) error {
	if err := r.Identifier(&c.Channel); err != nil {
		return err
	}
	return r.ReadAll(&c.Data)
}

type ServerboundPluginMessage struct {
	ClientboundPluginMessage
}

func (ServerboundPluginMessage) ID() int32 {
	return 0x02
}
