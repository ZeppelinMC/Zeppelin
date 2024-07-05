package login

import (
	"aether/net/io"

	"github.com/google/uuid"
)

type LoginStart struct {
	Name       string
	PlayerUUID uuid.UUID
}

func (LoginStart) ID() int32 {
	return 0x00
}

func (l *LoginStart) Encode(w io.Writer) error {
	if err := w.String(l.Name); err != nil {
		return err
	}
	return w.UUID(l.PlayerUUID)
}

func (l *LoginStart) Decode(r io.Reader) error {
	if err := r.String(&l.Name); err != nil {
		return err
	}
	return r.UUID(&l.PlayerUUID)
}
