package play

import "github.com/zeppelinmc/zeppelin/protocol/net/io"

// clientbound
const PacketIdSignedChatCommand = 0x05

type SignedChatCommand struct {
	Command         string
	Timestamp, Salt int64
	Arguments       []SignedArgument
	MessageCount    int32
	Acknowledged    io.FixedBitSet
}

type SignedArgument struct {
	Name      string
	Signature [256]byte
}

func (SignedChatCommand) ID() int32 {
	return PacketIdChatCommand
}

func (c *SignedChatCommand) Encode(w io.Writer) error {
	if err := w.String(c.Command); err != nil {
		return err
	}
	if err := w.Long(c.Timestamp); err != nil {
		return err
	}
	if err := w.Long(c.Salt); err != nil {
		return err
	}
	if err := w.VarInt(int32(len(c.Arguments))); err != nil {
		return err
	}
	for _, arg := range c.Arguments {
		if err := w.String(arg.Name); err != nil {
			return err
		}
		if err := w.FixedByteArray(arg.Signature[:]); err != nil {
			return err
		}
	}
	if err := w.VarInt(c.MessageCount); err != nil {
		return err
	}
	return w.FixedBitSet(c.Acknowledged)
}

func (c *SignedChatCommand) Decode(r io.Reader) error {
	if err := r.String(&c.Command); err != nil {
		return err
	}
	if err := r.Long(&c.Timestamp); err != nil {
		return err
	}
	if err := r.Long(&c.Salt); err != nil {
		return err
	}
	var length int32
	if _, err := r.VarInt(&length); err != nil {
		return err
	}
	c.Arguments = make([]SignedArgument, length)

	for _, arg := range c.Arguments {
		if err := r.String(&arg.Name); err != nil {
			return err
		}
		if err := r.FixedByteArray(arg.Signature[:]); err != nil {
			return err
		}
	}
	if _, err := r.VarInt(&c.MessageCount); err != nil {
		return err
	}
	return r.FixedBitSet(&c.Acknowledged, 20)
}
