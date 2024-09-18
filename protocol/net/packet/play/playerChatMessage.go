package play

import (
	"github.com/google/uuid"
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
	"github.com/zeppelinmc/zeppelin/protocol/text"
)

const (
	FilterTypePassThrough = iota
	FilterTypeFullyFiltered
	FilterTypePartiallyFiltered
)

type PreviousMessage struct {
	MessageID int32
	Signature *[256]byte
}

// clientbound
const PacketIdPlayerChatMessage = 0x39

type PlayerChatMessage struct {
	Sender              uuid.UUID
	Index               int32
	HasMessageSignature bool
	MessageSignature    [256]byte

	Message         string
	Timestamp, Salt int64

	PreviousMessages []PreviousMessage

	UnsignedContent *text.TextComponent
	FilterType      int32
	FilterBits      encoding.BitSet

	ChatType   int32
	SenderName text.TextComponent

	TargetName *text.TextComponent
}

func (PlayerChatMessage) ID() int32 {
	return PacketIdPlayerChatMessage
}

func (p *PlayerChatMessage) Encode(w encoding.Writer) error {
	if err := w.UUID(p.Sender); err != nil {
		return err
	}
	if err := w.VarInt(p.Index); err != nil {
		return err
	}
	if err := w.Bool(p.HasMessageSignature); err != nil {
		return err
	}
	if p.HasMessageSignature {
		if err := w.FixedByteArray(p.MessageSignature[:]); err != nil {
			return err
		}
	}

	if err := w.String(p.Message); err != nil {
		return err
	}
	if err := w.Long(p.Timestamp); err != nil {
		return err
	}
	if err := w.Long(p.Salt); err != nil {
		return err
	}

	if err := w.VarInt(int32(len(p.PreviousMessages))); err != nil {
		return err
	}
	for _, sig := range p.PreviousMessages {
		if err := w.VarInt(sig.MessageID + 1); err != nil {
			return err
		}
		if sig.MessageID+1 == 0 {
			if err := w.FixedByteArray(sig.Signature[:]); err != nil {
				return err
			}
		}
	}

	if err := w.Bool(p.UnsignedContent != nil); err != nil {
		return err
	}
	if p.UnsignedContent != nil {
		if err := w.TextComponent(*p.UnsignedContent); err != nil {
			return err
		}
	}
	if err := w.VarInt(p.FilterType); err != nil {
		return err
	}
	if p.FilterType == FilterTypePartiallyFiltered {
		if err := w.BitSet(p.FilterBits); err != nil {
			return err
		}
	}

	if err := w.VarInt(p.ChatType); err != nil {
		return err
	}
	if err := w.TextComponent(p.SenderName); err != nil {
		return err
	}

	if err := w.Bool(p.TargetName != nil); err != nil {
		return err
	}
	if p.TargetName != nil {
		if err := w.TextComponent(*p.TargetName); err != nil {
			return err
		}
	}

	return nil
}

func (p *PlayerChatMessage) Decode(r encoding.Reader) error {
	return nil //TODO
}
