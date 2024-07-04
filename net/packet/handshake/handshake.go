package handshake

import (
	"aether/net/io"
)

const (
	Status = iota + 1
	Login
	Transfer
)

type Handshaking struct {
	ProtocolVersion int32
	ServerAddress   string
	ServerPort      uint16
	NextState       int32
}

func (Handshaking) ID() int32 {
	return 0x00
}

func (h *Handshaking) Decode(r io.Reader) error {
	if _, err := r.VarInt(&h.ProtocolVersion); err != nil {
		return err
	}
	if err := r.String(&h.ServerAddress); err != nil {
		return err
	}
	if err := r.Ushort(&h.ServerPort); err != nil {
		return err
	}
	_, err := r.VarInt(&h.NextState)
	return err
}

func (h Handshaking) Encode(w io.Writer) error {
	if err := w.VarInt(h.ProtocolVersion); err != nil {
		return err
	}
	if err := w.String(h.ServerAddress); err != nil {
		return err
	}
	if err := w.Ushort(h.ServerPort); err != nil {
		return err
	}
	return w.VarInt(h.NextState)
}
