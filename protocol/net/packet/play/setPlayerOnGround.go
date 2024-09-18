package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// serverbound
const PacketIdSetPlayerOnGround = 0x1D

type SetPlayerOnGround struct {
	OnGround bool
}

func (SetPlayerOnGround) ID() int32 {
	return PacketIdSetPlayerOnGround
}

func (s *SetPlayerOnGround) Encode(w encoding.Writer) error {
	return w.Bool(s.OnGround)
}

func (s *SetPlayerOnGround) Decode(r encoding.Reader) error {
	return r.Bool(&s.OnGround)
}
