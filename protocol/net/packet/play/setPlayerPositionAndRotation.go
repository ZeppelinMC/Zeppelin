package play

import "github.com/zeppelinmc/zeppelin/protocol/net/io"

// serverbound
const PacketIdSetPlayerPositionAndRotation = 0x1B

type SetPlayerPositionAndRotation struct {
	X, Y, Z    float64
	Yaw, Pitch float32
	OnGround   bool
}

func (SetPlayerPositionAndRotation) ID() int32 {
	return 0x1B
}

func (s *SetPlayerPositionAndRotation) Encode(w io.Writer) error {
	if err := w.Double(s.X); err != nil {
		return err
	}
	if err := w.Double(s.Y); err != nil {
		return err
	}
	if err := w.Double(s.Z); err != nil {
		return err
	}
	if err := w.Float(s.Yaw); err != nil {
		return err
	}
	if err := w.Float(s.Pitch); err != nil {
		return err
	}
	return w.Bool(s.OnGround)
}

func (s *SetPlayerPositionAndRotation) Decode(r io.Reader) error {
	if err := r.Double(&s.X); err != nil {
		return err
	}
	if err := r.Double(&s.Y); err != nil {
		return err
	}
	if err := r.Double(&s.Z); err != nil {
		return err
	}
	if err := r.Float(&s.Yaw); err != nil {
		return err
	}
	if err := r.Float(&s.Pitch); err != nil {
		return err
	}
	return r.Bool(&s.OnGround)
}
