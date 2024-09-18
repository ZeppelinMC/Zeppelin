package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// clientbound
const PacketIdEntitySoundEffect = 0x67

type EntitySoundEffect struct {
	SoundId       int32   // -1 for custom
	SoundName     string  // only if sound id -1
	FixedRange    bool    // only if sound id -1
	Range         float32 // only if fixed range
	SoundCategory int32   // one of the above constants
	EntityId      int32
	Volume        float32
	Pitch         float32
	Seed          int64
}

func (EntitySoundEffect) ID() int32 {
	return PacketIdEntitySoundEffect
}

func (s *EntitySoundEffect) Encode(w encoding.Writer) error {
	if err := w.VarInt(s.SoundId + 1); err != nil {
		return err
	}
	if s.SoundId == -1 {
		if err := w.Identifier(s.SoundName); err != nil {
			return err
		}
		if err := w.Bool(s.FixedRange); err != nil {
			return err
		}
		if s.FixedRange {
			if err := w.Float(s.Range); err != nil {
				return err
			}
		}
	}
	if err := w.VarInt(s.SoundCategory); err != nil {
		return err
	}
	if err := w.VarInt(s.EntityId); err != nil {
		return err
	}
	if err := w.Float(s.Volume); err != nil {
		return err
	}
	if err := w.Float(s.Pitch); err != nil {
		return err
	}
	return w.Long(s.Seed)
}

func (s *EntitySoundEffect) Decode(r encoding.Reader) error {
	return nil //TODO
}
