package login

import "aether/net/io"

// clientbound
const PacketIdCookieRequest = 0x05

type CookieRequest struct {
	Key string
}

func (CookieRequest) ID() int32 {
	return 0x05
}

func (s *CookieRequest) Encode(w io.Writer) error {
	return w.String(s.Key)
}

func (s *CookieRequest) Decode(r io.Reader) error {
	return r.String(&s.Key)
}

type CookieResponse struct {
	Key     string
	Found   bool
	Payload []byte
}

func (CookieResponse) ID() int32 {
	return 0x04
}

func (s *CookieResponse) Encode(w io.Writer) error {
	if err := w.String(s.Key); err != nil {
		return err
	}
	if err := w.Bool(s.Found); err != nil {
		return err
	}
	if s.Found {
		return w.ByteArray(s.Payload)
	}
	return nil
}

func (s *CookieResponse) Decode(r io.Reader) error {
	if err := r.String(&s.Key); err != nil {
		return err
	}
	if err := r.Bool(&s.Found); err != nil {
		return err
	}
	if s.Found {
		return r.ByteArray(&s.Payload)
	}
	return nil
}
