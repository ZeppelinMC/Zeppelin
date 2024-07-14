package play

import (
	"github.com/dynamitemc/aether/net/io"
	"github.com/google/uuid"
)

const PacketIdPlayerInfoRemove = 0x3D

type PlayerInfoRemove struct {
	UUIDs []uuid.UUID
}

func (PlayerInfoRemove) ID() int32 {
	return 0x3D
}

func (p *PlayerInfoRemove) Encode(w io.Writer) error {
	if err := w.VarInt(int32(len(p.UUIDs))); err != nil {
		return err
	}
	for _, uuid := range p.UUIDs {
		if err := w.UUID(uuid); err != nil {
			return err
		}
	}
	return nil
}

func (p *PlayerInfoRemove) Decode(r io.Reader) error {
	var length int32
	if _, err := r.VarInt(&length); err != nil {
		return err
	}
	p.UUIDs = make([]uuid.UUID, length)
	for _, uuid := range p.UUIDs {
		if err := r.UUID(&uuid); err != nil {
			return err
		}
	}
	return nil
}
