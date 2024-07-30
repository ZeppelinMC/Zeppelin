package play

import "github.com/zeppelinmc/zeppelin/net/io"

// serverbound
const PacketIdSetPlayerOnGround = 0x1D

type SetPlayerOnGround struct {
	OnGround bool
}

func (SetPlayerOnGround) ID() int32 {
	return PacketIdSetPlayerOnGround
}

func (s *SetPlayerOnGround) Encode(w io.Writer) error {
	return w.Bool(s.OnGround)
}

func (s *SetPlayerOnGround) Decode(r io.Reader) error {
	return r.Bool(&s.OnGround)
}
