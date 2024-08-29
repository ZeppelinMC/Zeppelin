package play

import "github.com/zeppelinmc/zeppelin/protocol/net/io"

// clientbound
const PacketIdBlockUpdate = 0x09

type BlockUpdate struct {
	X, Y, Z int32
	BlockId int32
}

func (BlockUpdate) ID() int32 {
	return PacketIdBlockUpdate
}

func (b *BlockUpdate) Encode(w io.Writer) error {
	if err := w.Position(b.X, b.Y, b.Z); err != nil {
		return err
	}
	return w.VarInt(b.BlockId)
}

func (b *BlockUpdate) Decode(r io.Reader) error {
	if err := r.Position(&b.X, &b.Y, &b.Z); err != nil {
		return err
	}
	_, err := r.VarInt(&b.BlockId)
	return err
}
