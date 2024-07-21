package play

import "github.com/zeppelinmc/zeppelin/net/io"

// clientbound
const PacketIdSetDefaultSpawnPosition = 0x56

type SetDefaultSpawnPosition struct {
	X, Y, Z int32
	Angle   float32
}

func (SetDefaultSpawnPosition) ID() int32 {
	return PacketIdSetDefaultSpawnPosition
}

func (s *SetDefaultSpawnPosition) Encode(w io.Writer) error {
	if err := w.Position(s.X, s.Y, s.Z); err != nil {
		return err
	}
	return w.Float(s.Angle)
}

func (s *SetDefaultSpawnPosition) Decode(r io.Reader) error {
	if err := r.Position(&s.X, &s.Y, &s.Z); err != nil {
		return err
	}
	return r.Float(&s.Angle)
}
