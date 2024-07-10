package play

import "github.com/dynamitemc/aether/net/io"

//clientbound
const PacketIdSetCenterChunk = 0x54

type SetCenterChunk struct {
	ChunkX, ChunkZ int32
}

func (SetCenterChunk) ID() int32 {
	return 0x54
}

func (s *SetCenterChunk) Encode(w io.Writer) error {
	if err := w.VarInt(s.ChunkX); err != nil {
		return err
	}
	return w.VarInt(s.ChunkZ)
}

func (s *SetCenterChunk) Decode(r io.Reader) error {
	if _, err := r.VarInt(&s.ChunkX); err != nil {
		return err
	}
	_, err := r.VarInt(&s.ChunkZ)
	return err
}
