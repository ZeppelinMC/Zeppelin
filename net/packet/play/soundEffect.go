package play

import "github.com/zeppelinmc/zeppelin/net/io"

// clientbound
const PacketIdSoundEffect = 0x68

const (
	SoundCategoryMaster = iota
	SoundCategoryMusic
	SoundCategoryRecord
	SoundCategoryWeather
	SoundCategoryBlock
	SoundCategoryHostile
	SoundCategoryNeutral
	SoundCategoryPlayer
	SoundCategoryAmbient
	SoundCategoryVoice
)

type SoundEffect struct {
	SoundId       int32   // -1 for custom
	SoundName     string  // only if sound id -1
	FixedRange    bool    // only if sound id -1
	Range         float32 // only if fixed range
	SoundCategory int32   // one of the above constants
	X, Y, Z       int32   // the original position of this sound. Calculations are done on encoding
	Volume        float32
	Pitch         float32
	Seed          int64
}

func (SoundEffect) ID() int32 {
	return PacketIdSoundEffect
}

func (s *SoundEffect) Encode(w io.Writer) error {
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
	if err := w.Int(s.X * 8); err != nil {
		return err
	}
	if err := w.Int(s.Y * 8); err != nil {
		return err
	}
	if err := w.Int(s.Z * 8); err != nil {
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

func (s *SoundEffect) Decode(r io.Reader) error {
	return nil //TODO
}
