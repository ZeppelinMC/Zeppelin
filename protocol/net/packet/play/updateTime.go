package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// clientbound
const PacketIdUpdateTime = 0x64

type UpdateTime struct {
	WorldAge  int64
	TimeOfDay int64
}

func (UpdateTime) ID() int32 {
	return PacketIdUpdateTime
}

func (u *UpdateTime) Encode(w encoding.Writer) error {
	if err := w.Long(u.WorldAge); err != nil {
		return err
	}
	return w.Long(u.TimeOfDay)
}

func (u *UpdateTime) Decode(r encoding.Reader) error {
	if err := r.Long(&u.WorldAge); err != nil {
		return err
	}
	return r.Long(&u.TimeOfDay)
}
