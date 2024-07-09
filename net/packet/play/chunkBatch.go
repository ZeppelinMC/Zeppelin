package play

import (
	"aether/net/io"
	"aether/net/packet"
)

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

type ChunkBatchStart packet.EmptyPacket

func (ChunkBatchStart) ID() int32 {
	return 0x0D
}
