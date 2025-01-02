package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// serverbound
const PacketIdChatMessage = 0x06

type ChatMessage struct {
	Message         string
	Timestamp, Salt int64

	HasSignature bool
	Signature    [256]byte

	MessageCount int32
	Acknowledged encoding.FixedBitSet
}

func (ChatMessage) ID() int32 {
	return 0x06
}

func (c *ChatMessage) Encode(w encoding.Writer) error {
	if err := w.String(c.Message); err != nil {
		return err
	}
	if err := w.Long(c.Timestamp); err != nil {
		return err
	}
	if err := w.Long(c.Salt); err != nil {
		return err
	}
	if err := w.Bool(c.HasSignature); err != nil {
		return err
	}
	if c.HasSignature {
		if err := w.FixedByteArray(c.Signature[:]); err != nil {
			return err
		}
	}
	if err := w.VarInt(c.MessageCount); err != nil {
		return err
	}
	return w.FixedBitSet(c.Acknowledged)
}

func (c *ChatMessage) Decode(r encoding.Reader) error {
	if err := r.String(&c.Message); err != nil {
		return err
	}
	if err := r.Long(&c.Timestamp); err != nil {
		return err
	}
	if err := r.Long(&c.Salt); err != nil {
		return err
	}
	if err := r.Bool(&c.HasSignature); err != nil {
		return err
	}
	if c.HasSignature {
		if err := r.FixedByteArray(c.Signature[:]); err != nil {
			return err
		}
	}
	if _, err := r.VarInt(&c.MessageCount); err != nil {
		return err
	}
	return r.FixedBitSet(&c.Acknowledged, 20)
}
