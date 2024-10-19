package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
	"github.com/zeppelinmc/zeppelin/protocol/net/slot"
)

// serverbound
const PacketIdClickContainer = 0x0E

type ClickContainer struct {
	WindowId     byte
	State        int32
	Slot         int16
	Button       int8
	Mode         int32
	ChangedSlots []ChangedSlot
	CarriedItem  slot.Slot
}

type ChangedSlot struct {
	slot.Slot
	Num int16
}

func (ClickContainer) ID() int32 {
	return PacketIdClickContainer
}

func (c *ClickContainer) Decode(r encoding.Reader) error {
	if err := r.Ubyte(&c.WindowId); err != nil {
		return err
	}
	if _, err := r.VarInt(&c.State); err != nil {
		return err
	}
	if err := r.Short(&c.Slot); err != nil {
		return err
	}
	if err := r.Byte(&c.Button); err != nil {
		return err
	}
	if _, err := r.VarInt(&c.Mode); err != nil {
		return err
	}
	var len int32
	if _, err := r.VarInt(&len); err != nil {
		return err
	}
	c.ChangedSlots = make([]ChangedSlot, len)

	for i := int32(0); i < len; i++ {
		if err := r.Short(&c.ChangedSlots[i].Num); err != nil {
			return err
		}
		if err := c.ChangedSlots[i].Decode(r); err != nil {
			return err
		}
	}

	return c.CarriedItem.Decode(r)
}
