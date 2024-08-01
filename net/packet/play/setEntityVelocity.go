package play

import "github.com/zeppelinmc/zeppelin/net/io"

// clientbound
const PacketIdSetEntityVelocity = 0x5A

type SetEntityVelocity struct {
	EntityId int32
	X, Y, Z  int16
}

func (SetEntityVelocity) ID() int32 {
	return PacketIdSetEntityVelocity
}

func (d *SetEntityVelocity) Encode(w io.Writer) error {
	if err := w.VarInt(d.EntityId); err != nil {
		return err
	}
	if err := w.Short(d.X); err != nil {
		return err
	}
	if err := w.Short(d.Y); err != nil {
		return err
	}
	return w.Short(d.Z)
}

func (d *SetEntityVelocity) Decode(r io.Reader) error {
	if _, err := r.VarInt(&d.EntityId); err != nil {
		return err
	}
	if err := r.Short(&d.X); err != nil {
		return err
	}
	if err := r.Short(&d.Y); err != nil {
		return err
	}
	return r.Short(&d.Z)
}
