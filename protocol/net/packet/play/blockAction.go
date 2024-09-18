package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// clientbound
const PacketIdBlockAction = 0x08

type BlockAction struct {
	X, Y, Z                   int32
	ActionId, ActionParameter byte
	BlockType                 int32
}

func (BlockAction) ID() int32 {
	return PacketIdBlockAction
}

func (b *BlockAction) Encode(w encoding.Writer) error {
	if err := w.Position(b.X, b.Y, b.Z); err != nil {
		return err
	}
	if err := w.Ubyte(b.ActionId); err != nil {
		return err
	}
	if err := w.Ubyte(b.ActionParameter); err != nil {
		return err
	}
	return w.VarInt(b.BlockType)
}

func (b *BlockAction) Decode(r encoding.Reader) error {
	if err := r.Position(&b.X, &b.Y, &b.Z); err != nil {
		return err
	}
	if err := r.Ubyte(&b.ActionId); err != nil {
		return err
	}
	if err := r.Ubyte(&b.ActionParameter); err != nil {
		return err
	}
	_, err := r.VarInt(&b.BlockType)
	return err
}
