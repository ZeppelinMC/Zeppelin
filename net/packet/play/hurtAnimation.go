package play

import "github.com/zeppelinmc/zeppelin/net/io"

// clientbound
const PacketIdHurtAnimation = 0x24

type HurtAnimation struct {
	EntityId int32
	Yaw      float32
}

func (HurtAnimation) ID() int32 {
	return PacketIdHurtAnimation
}

func (d *HurtAnimation) Encode(w io.Writer) error {
	if err := w.VarInt(d.EntityId); err != nil {
		return err
	}
	return w.Float(d.Yaw)
}

func (d *HurtAnimation) Decode(r io.Reader) error {
	if _, err := r.VarInt(&d.EntityId); err != nil {
		return err
	}
	return r.Float(&d.Yaw)
}
