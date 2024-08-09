package play

import "github.com/zeppelinmc/zeppelin/net/io"

// clientbound
const PacketIdSetTickingState = 0x71

type SetTickingState struct {
	TickRate float32
	IsFrozen bool
}

func (SetTickingState) ID() int32 {
	return PacketIdSetTickingState
}

func (s *SetTickingState) Encode(w io.Writer) error {
	if err := w.Float(s.TickRate); err != nil {
		return err
	}
	return w.Bool(s.IsFrozen)
}

func (s *SetTickingState) Decode(r io.Reader) error {
	if err := r.Float(&s.TickRate); err != nil {
		return err
	}
	return r.Bool(&s.IsFrozen)
}
