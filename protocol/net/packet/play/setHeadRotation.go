package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// clientbound
const PacketIdSetHeadRotation = 0x48

type SetHeadRotation struct {
	EntityId int32
	HeadYaw  byte
}

func (SetHeadRotation) ID() int32 {
	return PacketIdSetHeadRotation
}

func (s *SetHeadRotation) Encode(w encoding.Writer) error {
	if err := w.VarInt(s.EntityId); err != nil {
		return err
	}
	return w.Ubyte(s.HeadYaw)
}

func (s *SetHeadRotation) Decode(r encoding.Reader) error {
	if _, err := r.VarInt(&s.EntityId); err != nil {
		return err
	}
	return r.Ubyte(&s.HeadYaw)
}
