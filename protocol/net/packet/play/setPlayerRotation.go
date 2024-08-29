package play

import "github.com/zeppelinmc/zeppelin/protocol/net/io"

// serverbound
const PacketIdSetPlayerRotation = 0x1C

type SetPlayerRotation struct {
	Yaw, Pitch float32
	OnGround   bool
}

func (SetPlayerRotation) ID() int32 {
	return 0x1C
}

func (s *SetPlayerRotation) Encode(w io.Writer) error {
	if err := w.Float(s.Yaw); err != nil {
		return err
	}
	if err := w.Float(s.Pitch); err != nil {
		return err
	}
	return w.Bool(s.OnGround)
}

func (s *SetPlayerRotation) Decode(r io.Reader) error {
	if err := r.Float(&s.Yaw); err != nil {
		return err
	}
	if err := r.Float(&s.Pitch); err != nil {
		return err
	}
	return r.Bool(&s.OnGround)
}
