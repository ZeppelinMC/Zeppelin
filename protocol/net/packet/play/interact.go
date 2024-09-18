package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// serverbound
const PacketIdInteract = 0x16

const (
	InteractTypeInteract = iota
	InteractTypeAttack
	InteractTypeInteractAt
)

type Interact struct {
	EntityId                  int32
	Type                      int32
	TargetX, TargetY, TargetZ float32
	Hand                      int32
	Sneaking                  bool
}

func (Interact) ID() int32 {
	return PacketIdInteract
}

func (i *Interact) Encode(w encoding.Writer) error {
	if err := w.VarInt(i.EntityId); err != nil {
		return err
	}
	if err := w.VarInt(i.Type); err != nil {
		return err
	}
	if i.Type == InteractTypeInteractAt {
		if err := w.Float(i.TargetX); err != nil {
			return err
		}
		if err := w.Float(i.TargetY); err != nil {
			return err
		}
		if err := w.Float(i.TargetZ); err != nil {
			return err
		}
	}
	if i.Type == InteractTypeInteractAt || i.Type == InteractTypeInteract {
		if err := w.VarInt(i.Hand); err != nil {
			return err
		}
	}

	return w.Bool(i.Sneaking)
}

func (i *Interact) Decode(r encoding.Reader) error {
	if _, err := r.VarInt(&i.EntityId); err != nil {
		return err
	}
	if _, err := r.VarInt(&i.Type); err != nil {
		return err
	}
	if i.Type == InteractTypeInteractAt {
		if err := r.Float(&i.TargetX); err != nil {
			return err
		}
		if err := r.Float(&i.TargetY); err != nil {
			return err
		}
		if err := r.Float(&i.TargetZ); err != nil {
			return err
		}
	}
	if i.Type == InteractTypeInteractAt || i.Type == InteractTypeInteract {
		if _, err := r.VarInt(&i.Hand); err != nil {
			return err
		}
	}

	return r.Bool(&i.Sneaking)
}
