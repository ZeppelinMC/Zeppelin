package play

import "github.com/zeppelinmc/zeppelin/protocol/net/io"

// clientbound
const PacketIdChangeDifficulty = 0x0B

type ChangeDifficulty struct {
	Difficulty byte
	Locked     bool
}

func (ChangeDifficulty) ID() int32 {
	return PacketIdChangeDifficulty
}

func (c *ChangeDifficulty) Encode(w io.Writer) error {
	if err := w.Ubyte(c.Difficulty); err != nil {
		return err
	}
	return w.Bool(c.Locked)
}

func (c *ChangeDifficulty) Decode(r io.Reader) error {
	if err := r.Ubyte(&c.Difficulty); err != nil {
		return err
	}
	return r.Bool(&c.Locked)
}
