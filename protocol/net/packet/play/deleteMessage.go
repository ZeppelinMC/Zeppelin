package play

import "github.com/zeppelinmc/zeppelin/protocol/net/io"

// clientbound
const PacketIdDeleteMessage = 0x1C

type DeleteMessage struct {
	MessageId int32
	Signature [256]byte
}

func (DeleteMessage) ID() int32 {
	return PacketIdDeleteMessage
}

func (b *DeleteMessage) Encode(w io.Writer) error {
	if err := w.VarInt(b.MessageId + 1); err != nil {
		return err
	}
	if b.MessageId == -1 {
		return w.FixedByteArray(b.Signature[:])
	}
	return nil
}

func (b *DeleteMessage) Decode(r io.Reader) error {
	if _, err := r.VarInt(&b.MessageId); err != nil {
		return err
	}
	if b.MessageId == -1 {
		return r.FixedByteArray(b.Signature[:])
	}
	return nil
}
