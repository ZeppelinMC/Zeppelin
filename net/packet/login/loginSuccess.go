package login

import (
	"github.com/dynamitemc/aether/net/io"

	"github.com/google/uuid"
)

type Property struct {
	Name      string
	Value     string
	Signature string
}

// clientbound
const PacketIdLoginSuccess = 0x02

type LoginSuccess struct {
	UUID                uuid.UUID
	Username            string
	Properties          []Property
	StrictErrorHandling bool
}

func (LoginSuccess) ID() int32 {
	return 0x02
}

func (l *LoginSuccess) Encode(w io.Writer) error {
	if err := w.UUID(l.UUID); err != nil {
		return err
	}
	if err := w.String(l.Username); err != nil {
		return err
	}
	if err := w.VarInt(int32(len(l.Properties))); err != nil {
		return err
	}
	for _, property := range l.Properties {
		if err := w.String(property.Name); err != nil {
			return err
		}
		if err := w.String(property.Value); err != nil {
			return err
		}
		if err := w.Bool(property.Signature != ""); err != nil {
			return err
		}
		if property.Signature != "" {
			if err := w.String(property.Signature); err != nil {
				return err
			}
		}
	}
	return w.Bool(l.StrictErrorHandling)
}

func (l *LoginSuccess) Decode(r io.Reader) error {
	if err := r.UUID(&l.UUID); err != nil {
		return err
	}
	if err := r.String(&l.Username); err != nil {
		return err
	}
	var length int32
	if _, err := r.VarInt(&length); err != nil {
		return err
	}
	l.Properties = make([]Property, length)
	for _, property := range l.Properties {
		if err := r.String(&property.Name); err != nil {
			return err
		}
		if err := r.String(&property.Value); err != nil {
			return err
		}
		var signed bool
		if err := r.Bool(&signed); err != nil {
			return err
		}
		if signed {
			if err := r.String(&property.Signature); err != nil {
				return err
			}
		}
	}
	return r.Bool(&l.StrictErrorHandling)
}

const PacketIdLoginAcknowledged = 0x03

type LoginAcknowledged struct{}

func (LoginAcknowledged) ID() int32 {
	return 0x03
}

func (*LoginAcknowledged) Encode(io.Writer) error {
	return nil
}

func (*LoginAcknowledged) Decode(io.Reader) error {
	return nil
}
