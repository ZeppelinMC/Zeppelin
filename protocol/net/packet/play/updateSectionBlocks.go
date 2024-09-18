package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// clientbound
const PacketIdUpdateSectionBlocks = 0x49

type UpdateSectionBlocks struct {
	ChunkX, ChunkY, ChunkZ int32

	// [x, y, z] -> state id
	Blocks map[[3]int32]int32
}

func (UpdateSectionBlocks) ID() int32 {
	return PacketIdUpdateSectionBlocks
}

func (b *UpdateSectionBlocks) Encode(w encoding.Writer) error {
	if err := w.Long(((int64(b.ChunkX) & 0x3FFFFF) << 42) | (int64(b.ChunkY) & 0xFFFFF) | ((int64(b.ChunkZ) & 0x3FFFFF) << 20)); err != nil {
		return err
	}
	if err := w.VarInt(int32(len(b.Blocks))); err != nil {
		return err
	}
	for pos, state := range b.Blocks {
		blockLocalX, blockLocalY, blockLocalZ := int64(pos[0]), int64(pos[1]), int64(pos[2])
		if err := w.VarLong(int64(state)<<12 | (blockLocalX<<8 | blockLocalZ<<4 | blockLocalY)); err != nil {
			return err
		}
	}
	return nil
}

func (b *UpdateSectionBlocks) Decode(r encoding.Reader) error {
	var sectionPos int64
	if err := r.Long(&sectionPos); err != nil {
		return err
	}
	b.ChunkX = int32(sectionPos >> 42)
	b.ChunkY = int32(sectionPos << 44 >> 44)
	b.ChunkZ = int32(sectionPos << 22 >> 42)

	var blocksLen int32
	if _, err := r.VarInt(&blocksLen); err != nil {
		return err
	}
	b.Blocks = make(map[[3]int32]int32, blocksLen)

	for i := 0; i < int(blocksLen); i++ {
		var blockId int64
		if err := r.Long(&blockId); err != nil {
			return err
		}
		var pos = [3]int32{
			int32((blockId >> 8) & 0xF),
			int32(blockId & 0xF),
			int32((blockId >> 4) & 0xF),
		}

		b.Blocks[pos] = int32(blockId >> 12)
	}

	return nil
}
