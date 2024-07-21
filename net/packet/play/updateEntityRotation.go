package play

import "github.com/zeppelinmc/zeppelin/net/io"

// clientbound
const PacketIdUpdateEntityRotation = 0x30

type UpdateEntityRotation struct {
	EntityId   int32
	Yaw, Pitch byte
	OnGround   bool
}

func (UpdateEntityRotation) ID() int32 {
	return PacketIdUpdateEntityRotation
}

func (s *UpdateEntityRotation) Encode(w io.Writer) error {
	if err := w.VarInt(s.EntityId); err != nil {
		return err
	}
	if err := w.Ubyte(s.Yaw); err != nil {
		return err
	}
	if err := w.Ubyte(s.Pitch); err != nil {
		return err
	}
	return w.Bool(s.OnGround)
}

func (s *UpdateEntityRotation) Decode(r io.Reader) error {
	if _, err := r.VarInt(&s.EntityId); err != nil {
		return err
	}
	if err := r.Ubyte(&s.Yaw); err != nil {
		return err
	}
	if err := r.Ubyte(&s.Pitch); err != nil {
		return err
	}
	return r.Bool(&s.OnGround)
}
