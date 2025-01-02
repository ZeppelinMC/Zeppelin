package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet"
)

// clientbound
const PacketIdChunkBatchFinished = 0x0C

type ChunkBatchFinished struct {
	BatchSize int32
}

func (ChunkBatchFinished) ID() int32 {
	return PacketIdChunkBatchFinished
}

func (c *ChunkBatchFinished) Encode(w encoding.Writer) error {
	return w.VarInt(c.BatchSize)
}

func (c *ChunkBatchFinished) Decode(r encoding.Reader) error {
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

func (c *ChunkBatchReceived) Encode(w encoding.Writer) error {
	return w.Float(c.ChunksPerTick)
}

func (c *ChunkBatchReceived) Decode(r encoding.Reader) error {
	return r.Float(&c.ChunksPerTick)
}
