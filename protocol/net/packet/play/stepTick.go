package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// clientbound
const PacketIdStepTick = 0x72

type StepTick struct {
	TickSteps int32
}

func (StepTick) ID() int32 {
	return PacketIdStepTick
}

func (s *StepTick) Encode(w encoding.Writer) error {
	return w.VarInt(s.TickSteps)
}
