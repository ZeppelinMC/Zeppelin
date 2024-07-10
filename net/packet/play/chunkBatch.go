package play

import (
	"github.com/dynamitemc/aether/net/io"
	"github.com/dynamitemc/aether/net/packet"
)

// serverbound
const PacketIdChunkBatchFinished = 0x0D

type ChunkBatchFinished struct {
	BatchSize int32
}

func (ChunkBatchFinished) ID() int32 {
	return 0x0C
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
	return 0x0D
}
