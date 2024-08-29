package play

import "github.com/zeppelinmc/zeppelin/protocol/net/io"

// serverbound
const PacketIdSwingArm = 0x36

const (
	MainHand = iota
	Offhand
)

type SwingArm struct {
	Hand int32
}

func (SwingArm) ID() int32 {
	return PacketIdSwingArm
}

func (e *SwingArm) Encode(w io.Writer) error {
	return w.VarInt(e.Hand)
}

func (e *SwingArm) Decode(r io.Reader) error {
	_, err := r.VarInt(&e.Hand)
	return err
}
