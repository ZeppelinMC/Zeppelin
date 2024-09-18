package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// clientbound
const PacketIdEntityAnimation = 0x03

const (
	AnimationSwingMainArm = iota
	_
	AnimationLeaveBed
	AnimationSwingOffhand
	AnimationCriticalEffect
	AnimationMagicCriticalEffect
)

type EntityAnimation struct {
	EntityId  int32
	Animation byte
}

func (EntityAnimation) ID() int32 {
	return PacketIdEntityAnimation
}

func (e *EntityAnimation) Encode(w encoding.Writer) error {
	if err := w.VarInt(e.EntityId); err != nil {
		return err
	}
	return w.Ubyte(e.Animation)
}

func (e *EntityAnimation) Decode(r encoding.Reader) error {
	if _, err := r.VarInt(&e.EntityId); err != nil {
		return err
	}
	return r.Ubyte(&e.Animation)
}
