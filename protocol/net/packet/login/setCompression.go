package login

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// clientbound
const PacketIdSetCompression = 0x03

type SetCompression struct {
	Threshold int32
}

func (SetCompression) ID() int32 {
	return 0x03
}

func (s *SetCompression) Encode(w encoding.Writer) error {
	return w.VarInt(s.Threshold)
}

func (s *SetCompression) Decode(r encoding.Reader) error {
	_, err := r.VarInt(&s.Threshold)
	return err
}
