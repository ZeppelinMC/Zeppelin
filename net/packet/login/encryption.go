package login

import "aether/net/io"

type EncryptionRequest struct {
	PublicKey          []byte
	VerifyToken        []byte
	ShouldAuthenticate bool
}

func (EncryptionRequest) ID() int32 {
	return 0x01
}

func (e *EncryptionRequest) Encode(w io.Writer) error {
	if err := w.String(""); err != nil {
		return err
	}
	if err := w.ByteArray(e.PublicKey); err != nil {
		return err
	}
	if err := w.ByteArray(e.VerifyToken); err != nil {
		return err
	}
	return w.Bool(e.ShouldAuthenticate)
}

func (e *EncryptionRequest) Decode(r io.Reader) error {
	var s string
	if err := r.String(&s); err != nil {
		return err
	}
	if err := r.ByteArray(&e.PublicKey); err != nil {
		return err
	}
	if err := r.ByteArray(&e.VerifyToken); err != nil {
		return err
	}
	return r.Bool(&e.ShouldAuthenticate)
}

type EncryptionResponse struct {
	SharedSecret []byte
	VerifyToken  []byte
}

func (EncryptionResponse) ID() int32 {
	return 0x01
}

func (e *EncryptionResponse) Encode(w io.Writer) error {
	if err := w.ByteArray(e.SharedSecret); err != nil {
		return err
	}
	return w.ByteArray(e.VerifyToken)
}

func (e *EncryptionResponse) Decode(r io.Reader) error {
	if err := r.ByteArray(&e.SharedSecret); err != nil {
		return err
	}
	return r.ByteArray(&e.VerifyToken)
}
