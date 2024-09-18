package play

import (
	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// serverbound
const PacketIdPlayerSession = 0x07

type PlayerSession struct {
	SessionID    uuid.UUID
	ExpiresAt    int64
	PublicKey    []byte
	KeySignature []byte
}

func (PlayerSession) ID() int32 {
	return 0x07
}

func (p *PlayerSession) Encode(w encoding.Writer) error {
	if err := w.UUID(p.SessionID); err != nil {
		return err
	}
	if err := w.Long(p.ExpiresAt); err != nil {
		return err
	}
	if err := w.ByteArray(p.PublicKey); err != nil {
		return err
	}
	return w.ByteArray(p.KeySignature)
}

func (p *PlayerSession) Decode(r encoding.Reader) error {
	if err := r.UUID(&p.SessionID); err != nil {
		return err
	}
	if err := r.Long(&p.ExpiresAt); err != nil {
		return err
	}
	if err := r.ByteArray(&p.PublicKey); err != nil {
		return err
	}
	return r.ByteArray(&p.KeySignature)
}
