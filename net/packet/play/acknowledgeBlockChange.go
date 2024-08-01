package play

import "github.com/zeppelinmc/zeppelin/net/io"

// clientbound
const PacketIdAcknowledgeBlockChange = 0x05

type AcknowledgeBlockChange struct {
	SequenceId int32
}

func (AcknowledgeBlockChange) ID() int32 {
	return PacketIdAcknowledgeBlockChange
}

func (b *AcknowledgeBlockChange) Encode(w io.Writer) error {
	return w.VarInt(b.SequenceId)
}

func (b *AcknowledgeBlockChange) Decode(r io.Reader) error {
	_, err := r.VarInt(&b.SequenceId)
	return err
}
