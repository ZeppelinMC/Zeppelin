package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// clientbound
const PacketIdSetTickingState = 0x71

type SetTickingState struct {
	TickRate float32
	IsFrozen bool
}

func (SetTickingState) ID() int32 {
	return PacketIdSetTickingState
}

func (s *SetTickingState) Encode(w encoding.Writer) error {
	if err := w.Float(s.TickRate); err != nil {
		return err
	}
	return w.Bool(s.IsFrozen)
}
