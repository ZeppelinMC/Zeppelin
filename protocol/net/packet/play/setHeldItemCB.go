package play

import "github.com/zeppelinmc/zeppelin/protocol/net/io"

// clientbound
const PacketIdSetHeldItemClientbound = 0x53

type SetHeldItemClientbound struct {
	Slot int8
}

func (SetHeldItemClientbound) ID() int32 {
	return PacketIdSetHeldItemClientbound
}

func (s *SetHeldItemClientbound) Encode(w io.Writer) error {
	return w.Byte(s.Slot)
}

func (s *SetHeldItemClientbound) Decode(r io.Reader) error {
	return r.Byte(&s.Slot)
}
