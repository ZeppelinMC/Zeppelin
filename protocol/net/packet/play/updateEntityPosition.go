package play

import "github.com/zeppelinmc/zeppelin/protocol/net/io"

// clientbound
const PacketIdUpdateEntityPosition = 0x2E

type UpdateEntityPosition struct {
	EntityId               int32
	DeltaX, DeltaY, DeltaZ int16
	OnGround               bool
}

func (UpdateEntityPosition) ID() int32 {
	return PacketIdUpdateEntityPosition
}

func (s *UpdateEntityPosition) Encode(w io.Writer) error {
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
	return w.Bool(s.OnGround)
}

func (s *UpdateEntityPosition) Decode(r io.Reader) error {
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
	return r.Bool(&s.OnGround)
}
