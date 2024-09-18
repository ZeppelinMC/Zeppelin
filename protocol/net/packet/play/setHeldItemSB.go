package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// serverbound
const PacketIdSetHeldItemServerbound = 0x2F

type SetHeldItemServerbound struct {
	Slot int16
}

func (SetHeldItemServerbound) ID() int32 {
	return PacketIdSetHeldItemServerbound
}

func (s *SetHeldItemServerbound) Encode(w encoding.Writer) error {
	return w.Short(s.Slot)
}

func (s *SetHeldItemServerbound) Decode(r encoding.Reader) error {
	return r.Short(&s.Slot)
}
