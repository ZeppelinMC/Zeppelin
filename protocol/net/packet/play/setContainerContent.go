package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
	"github.com/zeppelinmc/zeppelin/protocol/net/slot"
)

// clientbound
const PacketIdSetContainerContent = 0x13

type SetContainerContent struct {
	WindowID    byte
	StateId     int32
	Slots       []slot.Slot
	CarriedItem slot.Slot
}

func (SetContainerContent) ID() int32 {
	return PacketIdSetContainerContent
}

func (s *SetContainerContent) Encode(w encoding.Writer) error {
	if err := w.Ubyte(s.WindowID); err != nil {
		return err
	}
	if err := w.VarInt(s.StateId); err != nil {
		return err
	}
	if err := w.VarInt(int32(len(s.Slots))); err != nil {
		return err
	}
	for _, slot := range s.Slots {
		if err := slot.Encode(w); err != nil {
			return err
		}
	}

	return s.CarriedItem.Encode(w)
}

func (s *SetContainerContent) Decode(r encoding.Reader) error {
	return nil //TODO
}
