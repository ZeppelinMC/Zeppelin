package play

import (
	"github.com/zeppelinmc/zeppelin/net/io"
)

type BlockEntity struct {
	X, Y, Z int32
	Type    int32
	Data    any
}

// clientbound
const PacketIdChunkDataUpdateLight = 0x27

type Heightmaps struct {
	MOTION_BLOCKING, WORLD_SURFACE [37]int64
}

type ChunkDataUpdateLight struct {
	CX, CZ                                                               int32
	Heightmaps                                                           Heightmaps
	Data                                                                 []byte
	BlockEntities                                                        []BlockEntity
	SkyLightMask, BlockLightMask, EmptySkyLightMask, EmptyBlockLightMask io.BitSet
	SkyLightArrays                                                       [][]byte
	BlockLightArrays                                                     [][]byte
}

func (ChunkDataUpdateLight) ID() int32 {
	return 0x27
}

func (c *ChunkDataUpdateLight) Encode(w io.Writer) error {
	if err := w.Int(c.CX); err != nil {
		return err
	}
	if err := w.Int(c.CZ); err != nil {
		return err
	}
	if err := w.NBT(c.Heightmaps); err != nil {
		return err
	}
	if err := w.ByteArray(c.Data); err != nil {
		return err
	}
	if err := w.VarInt(int32(len(c.BlockEntities))); err != nil {
		return err
	}
	for _, blockEntity := range c.BlockEntities {
		if err := w.Ubyte(((byte(blockEntity.X) & 0x0f) << 4) | (byte(blockEntity.Z) & 0x0f)); err != nil {
			return err
		}
		if err := w.Short(int16(blockEntity.Y)); err != nil {
			return err
		}
		if err := w.VarInt(blockEntity.Type); err != nil {
			return err
		}
		if err := w.NBT(blockEntity.Data); err != nil {
			return err
		}
	}

	if err := w.BitSet(c.SkyLightMask); err != nil {
		return err
	}
	if err := w.BitSet(c.BlockLightMask); err != nil {
		return err
	}
	if err := w.BitSet(c.EmptySkyLightMask); err != nil {
		return err
	}
	if err := w.BitSet(c.EmptyBlockLightMask); err != nil {
		return err
	}

	if err := w.VarInt(int32(len(c.SkyLightArrays))); err != nil {
		return err
	}
	for _, array := range c.SkyLightArrays {
		if err := w.ByteArray(array); err != nil {
			return err
		}
	}

	if err := w.VarInt(int32(len(c.BlockLightArrays))); err != nil {
		return err
	}
	for _, array := range c.BlockLightArrays {
		if err := w.ByteArray(array); err != nil {
			return err
		}
	}
	return nil
}

func (c *ChunkDataUpdateLight) Decode(io.Reader) error {
	return nil
}
