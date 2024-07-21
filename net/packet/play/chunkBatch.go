package play

import (
	"github.com/zeppelinmc/zeppelin/net/io"
	"github.com/zeppelinmc/zeppelin/net/packet"
)

// clientbound
const PacketIdChunkBatchFinished = 0x0D

type ChunkBatchFinished struct {
	BatchSize int32
}

func (ChunkBatchFinished) ID() int32 {
	return PacketIdChunkBatchFinished
}

func (c *ChunkBatchFinished) Encode(w io.Writer) error {
	return w.VarInt(c.BatchSize)
}

func (c *ChunkBatchFinished) Decode(r io.Reader) error {
	_, err := r.VarInt(&c.BatchSize)
	return err
}

// clientbound
const PacketIdChunkBatchStart = 0x0D

type ChunkBatchStart struct{ packet.EmptyPacket }

func (ChunkBatchStart) ID() int32 {
	return PacketIdChunkBatchStart
}

// serverbound
const PacketIdChunkBatchReceived = 0x08

type ChunkBatchReceived struct {
	ChunksPerTick float32
}

func (ChunkBatchReceived) ID() int32 {
	return PacketIdChunkBatchReceived
}

func (c *ChunkBatchReceived) Encode(w io.Writer) error {
	return w.Float(c.ChunksPerTick)
}

func (c *ChunkBatchReceived) Decode(r io.Reader) error {
	return r.Float(&c.ChunksPerTick)
}
