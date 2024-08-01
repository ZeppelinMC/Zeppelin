package play

import (
	"github.com/zeppelinmc/zeppelin/nbt"
	"github.com/zeppelinmc/zeppelin/net/io"
)

// clientbound
const PacketIdBlockEntityData = 0x07

type BlockEntityData struct {
	X, Y, Z int32
	Type    int32
	Data    any
}

func (BlockEntityData) ID() int32 {
	return PacketIdBlockEntityData
}

func (b *BlockEntityData) Encode(w io.Writer) error {
	if err := w.Position(b.X, b.Y, b.Z); err != nil {
		return err
	}
	if err := w.VarInt(b.Type); err != nil {
		return err
	}
	if b.Data == nil {
		return w.Byte(nbt.End)
	}
	return w.NBT(b.Data)
}

func (b *BlockEntityData) Decode(r io.Reader) error {
	return nil //TODO
}
