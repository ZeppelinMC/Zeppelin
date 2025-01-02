package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// clientbound
const PacketIdEntityEvent = 0x1F

type EntityEvent struct {
	EntityId     int32
	EntityStatus int8
}

func (EntityEvent) ID() int32 {
	return PacketIdEntityEvent
}

func (e *EntityEvent) Encode(w encoding.Writer) error {
	if err := w.Int(e.EntityId); err != nil {
		return err
	}
	return w.Byte(e.EntityStatus)
}

func (e *EntityEvent) Decode(r encoding.Reader) error {
	if err := r.Int(&e.EntityId); err != nil {
		return err
	}
	return r.Byte(&e.EntityStatus)
}
