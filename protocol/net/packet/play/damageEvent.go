package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// clientbound
const PacketIdDamageEvent = 0x1A

type DamageEvent struct {
	EntityId       int32
	SourceTypeId   int32
	SourceCauseId  int32 //-1 for none
	SourceDirectId int32 //-1 for none

	HasSourcePosition         bool
	SourceX, SourceY, SourceZ float64
}

func (DamageEvent) ID() int32 {
	return PacketIdDamageEvent
}

func (d *DamageEvent) Encode(w encoding.Writer) error {
	if err := w.VarInt(d.EntityId); err != nil {
		return err
	}
	if err := w.VarInt(d.SourceTypeId); err != nil {
		return err
	}
	if err := w.VarInt(d.SourceDirectId + 1); err != nil {
		return err
	}
	if err := w.VarInt(d.SourceCauseId + 1); err != nil {
		return err
	}
	if err := w.Bool(d.HasSourcePosition); err != nil {
		return err
	}
	if d.HasSourcePosition {
		if err := w.Double(d.SourceX); err != nil {
			return err
		}
		if err := w.Double(d.SourceY); err != nil {
			return err
		}
		if err := w.Double(d.SourceZ); err != nil {
			return err
		}
	}
	return nil
}

func (d *DamageEvent) Decode(r encoding.Reader) error {
	return nil //TODO
}
