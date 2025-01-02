package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// clientbound
const PacketIdAcknowledgeBlockChange = 0x05

type AcknowledgeBlockChange struct {
	SequenceId int32
}

func (AcknowledgeBlockChange) ID() int32 {
	return PacketIdAcknowledgeBlockChange
}

func (b *AcknowledgeBlockChange) Encode(w encoding.Writer) error {
	return w.VarInt(b.SequenceId)
}

func (b *AcknowledgeBlockChange) Decode(r encoding.Reader) error {
	_, err := r.VarInt(&b.SequenceId)
	return err
}
