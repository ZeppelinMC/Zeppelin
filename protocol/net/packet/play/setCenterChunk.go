package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// clientbound
const PacketIdSetCenterChunk = 0x54

type SetCenterChunk struct {
	ChunkX, ChunkZ int32
}

func (SetCenterChunk) ID() int32 {
	return 0x54
}

func (s *SetCenterChunk) Encode(w encoding.Writer) error {
	if err := w.VarInt(s.ChunkX); err != nil {
		return err
	}
	return w.VarInt(s.ChunkZ)
}

func (s *SetCenterChunk) Decode(r encoding.Reader) error {
	if _, err := r.VarInt(&s.ChunkX); err != nil {
		return err
	}
	_, err := r.VarInt(&s.ChunkZ)
	return err
}
