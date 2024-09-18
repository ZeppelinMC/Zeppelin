package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// clientbound
const PacketIdRemoveEntities = 0x42

type RemoveEntities struct {
	EntityIDs []int32
}

func (RemoveEntities) ID() int32 {
	return PacketIdRemoveEntities
}

func (r *RemoveEntities) Encode(w encoding.Writer) error {
	if err := w.VarInt(int32(len(r.EntityIDs))); err != nil {
		return err
	}

	for _, entityId := range r.EntityIDs {
		if err := w.VarInt(entityId); err != nil {
			return err
		}
	}

	return nil
}

func (e *RemoveEntities) Decode(r encoding.Reader) error {
	var length int32
	if _, err := r.VarInt(&length); err != nil {
		return err
	}

	e.EntityIDs = make([]int32, length)

	for _, entityId := range e.EntityIDs {
		if _, err := r.VarInt(&entityId); err != nil {
			return err
		}
	}

	return nil
}
