package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// clientbound
const PacketIdHurtAnimation = 0x24

type HurtAnimation struct {
	EntityId int32
	Yaw      float32
}

func (HurtAnimation) ID() int32 {
	return PacketIdHurtAnimation
}

func (d *HurtAnimation) Encode(w encoding.Writer) error {
	if err := w.VarInt(d.EntityId); err != nil {
		return err
	}
	return w.Float(d.Yaw)
}

func (d *HurtAnimation) Decode(r encoding.Reader) error {
	if _, err := r.VarInt(&d.EntityId); err != nil {
		return err
	}
	return r.Float(&d.Yaw)
}
