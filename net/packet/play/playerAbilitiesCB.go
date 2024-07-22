package play

import "github.com/zeppelinmc/zeppelin/net/io"

// clientbound
const PacketIdPlayerAbilitiesClientbound = 0x38

const (
	PlayerAbsInvulenrable = 1 << iota
	PlayerAbsFlying
	PlayerAbsCreativeMode
)

type PlayerAbilitiesClientbound struct {
	Flags       int8
	FlyingSpeed float32
	FOVModifier float32
}

func (PlayerAbilitiesClientbound) ID() int32 {
	return PacketIdPlayerAbilitiesClientbound
}

func (a *PlayerAbilitiesClientbound) Encode(w io.Writer) error {
	if err := w.Byte(a.Flags); err != nil {
		return err
	}
	if err := w.Float(a.FlyingSpeed); err != nil {
		return err
	}
	return w.Float(a.FOVModifier)
}

func (a *PlayerAbilitiesClientbound) Decode(r io.Reader) error {
	if err := r.Byte(&a.Flags); err != nil {
		return err
	}
	if err := r.Float(&a.FlyingSpeed); err != nil {
		return err
	}
	return r.Float(&a.FOVModifier)
}
