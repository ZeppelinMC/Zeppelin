package play

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/protocol/net/io"
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
	if length < 0 {
		return fmt.Errorf("negative length for make (player info remove decode)")
	}
	p.UUIDs = make([]uuid.UUID, length)
	for _, uuid := range p.UUIDs {
		if err := r.UUID(&uuid); err != nil {
			return err
		}
	}
	return nil
}
