package play

import "github.com/zeppelinmc/zeppelin/net/io"

// serverbound
const PacketIdConfirmTeleportation = 0x00

type ConfirmTeleportation struct {
	TeleportId int32
}

func (ConfirmTeleportation) ID() int32 {
	return PacketIdConfirmTeleportation
}

func (c *ConfirmTeleportation) Encode(w io.Writer) error {
	return w.VarInt(c.TeleportId)
}

func (c *ConfirmTeleportation) Decode(r io.Reader) error {
	_, err := r.VarInt(&c.TeleportId)
	return err
}
