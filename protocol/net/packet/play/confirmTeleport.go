package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// serverbound
const PacketIdConfirmTeleportation = 0x00

type ConfirmTeleportation struct {
	TeleportId int32
}

func (ConfirmTeleportation) ID() int32 {
	return PacketIdConfirmTeleportation
}

func (c *ConfirmTeleportation) Encode(w encoding.Writer) error {
	return w.VarInt(c.TeleportId)
}

func (c *ConfirmTeleportation) Decode(r encoding.Reader) error {
	_, err := r.VarInt(&c.TeleportId)
	return err
}
