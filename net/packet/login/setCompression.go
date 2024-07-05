package login

import "aether/net/io"

type SetCompression struct {
	Threshold int32
}

func (SetCompression) ID() int32 {
	return 0x03
}

func (s *SetCompression) Encode(w io.Writer) error {
	return w.VarInt(s.Threshold)
}

func (s *SetCompression) Decode(r io.Reader) error {
	_, err := r.VarInt(&s.Threshold)
	return err
}
