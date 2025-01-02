package play

import (
	"github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"
)

// serverbound
const PacketIdUseItemOn = 0x38

type UseItemOn struct {
	Hand                               int32
	BlockX, BlockY, BlockZ             int32
	Face                               int32
	CursorPosX, CursorPosY, CursorPosZ float32
	InsideBlock                        bool
	Sequence                           int32
}

const (
	FaceBottom = iota
	FaceTop
	FaceNorth
	FaceSouth
	FaceWest
	FaceEast
)

func (UseItemOn) ID() int32 {
	return PacketIdUseItemOn
}

func (u *UseItemOn) Encode(w encoding.Writer) error {
	if err := w.VarInt(u.Hand); err != nil {
		return err
	}
	if err := w.Position(u.BlockX, u.BlockY, u.BlockZ); err != nil {
		return err
	}
	if err := w.VarInt(u.Face); err != nil {
		return err
	}
	if err := w.Float(u.CursorPosX); err != nil {
		return err
	}
	if err := w.Float(u.CursorPosY); err != nil {
		return err
	}
	if err := w.Float(u.CursorPosZ); err != nil {
		return err
	}
	if err := w.Bool(u.InsideBlock); err != nil {
		return err
	}
	return w.VarInt(u.Sequence)
}

func (u *UseItemOn) Decode(r encoding.Reader) error {
	if _, err := r.VarInt(&u.Hand); err != nil {
		return err
	}
	if err := r.Position(&u.BlockX, &u.BlockY, &u.BlockZ); err != nil {
		return err
	}
	if _, err := r.VarInt(&u.Face); err != nil {
		return err
	}
	if err := r.Float(&u.CursorPosX); err != nil {
		return err
	}
	if err := r.Float(&u.CursorPosY); err != nil {
		return err
	}
	if err := r.Float(&u.CursorPosZ); err != nil {
		return err
	}
	if err := r.Bool(&u.InsideBlock); err != nil {
		return err
	}
	_, err := r.VarInt(&u.Sequence)
	return err
}
