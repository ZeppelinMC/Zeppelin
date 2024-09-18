package login

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// clientbound
const PacketIdLoginPluginRequest = 0x04

type LoginPluginRequest struct {
	MessageID int32
	Channel   string
	Data      []byte
}

func (LoginPluginRequest) ID() int32 {
	return 0x04
}

func (l *LoginPluginRequest) Encode(w encoding.Writer) error {
	if err := w.VarInt(l.MessageID); err != nil {
		return err
	}
	if err := w.Identifier(l.Channel); err != nil {
		return err
	}
	return w.FixedByteArray(l.Data)
}

func (l *LoginPluginRequest) Decode(r encoding.Reader) error {
	if _, err := r.VarInt(&l.MessageID); err != nil {
		return err
	}
	if err := r.Identifier(&l.Channel); err != nil {
		return err
	}
	return r.ReadAll(&l.Data)
}

// serverbound
const PacketIdLoginPluginResponse = 0x02

type LoginPluginResponse struct {
	MessageID int32
	Sucessful bool
	Data      []byte
}

func (LoginPluginResponse) ID() int32 {
	return 0x02
}

func (l *LoginPluginResponse) Encode(w encoding.Writer) error {
	if err := w.VarInt(l.MessageID); err != nil {
		return err
	}
	if err := w.Bool(l.Sucessful); err != nil {
		return err
	}
	if l.Sucessful {
		if err := w.FixedByteArray(l.Data); err != nil {
			return err
		}
	}
	return nil
}

func (l *LoginPluginResponse) Decode(r encoding.Reader) error {
	if _, err := r.VarInt(&l.MessageID); err != nil {
		return err
	}
	if err := r.Bool(&l.Sucessful); err != nil {
		return err
	}
	if l.Sucessful {
		if err := r.ReadAll(&l.Data); err != nil {
			return err
		}
	}
	return nil
}
