package play

import "github.com/zeppelinmc/zeppelin/net/io"

// clientbound
const PacketIdUpdateEntityPositionAndRotation = 0x2F

type UpdateEntityPositionAndRotation struct {
	EntityId               int32
	DeltaX, DeltaY, DeltaZ int16
	Yaw, Pitch             byte
	OnGround               bool
}

func (UpdateEntityPositionAndRotation) ID() int32 {
	return PacketIdUpdateEntityPositionAndRotation
}

func (s *UpdateEntityPositionAndRotation) Encode(w io.Writer) error {
	if err := w.VarInt(s.EntityId); err != nil {
		return err
	}
	if err := w.Short(s.DeltaX); err != nil {
		return err
	}
	if err := w.Short(s.DeltaY); err != nil {
		return err
	}
	if err := w.Short(s.DeltaZ); err != nil {
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

func (s *UpdateEntityPositionAndRotation) Decode(r io.Reader) error {
	if _, err := r.VarInt(&s.EntityId); err != nil {
		return err
	}
	if err := r.Short(&s.DeltaX); err != nil {
		return err
	}
	if err := r.Short(&s.DeltaY); err != nil {
		return err
	}
	if err := r.Short(&s.DeltaZ); err != nil {
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
