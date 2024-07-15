package play

import (
	"github.com/dynamitemc/aether/net/io"
	"github.com/google/uuid"
)

const (
	ObjectDataItemFrameDown = iota
	ObjectDataItemFrameUp
	ObjectDataNorth
	ObjectDataSouth
	ObjectDataWest
	ObjectDataEast
)

// clientbound
const PacketIdSpawnEntity = 0x01

type SpawnEntity struct {
	EntityId            int32
	EntityUUID          uuid.UUID
	Type                int32
	X, Y, Z             float64
	Pitch, Yaw, HeadYaw byte
	Data                int32
	VelX, VelY, VelZ    int16
}

func (SpawnEntity) ID() int32 {
	return 0x01
}

func (s *SpawnEntity) Encode(w io.Writer) error {
	if err := w.VarInt(s.EntityId); err != nil {
		return err
	}
	if err := w.UUID(s.EntityUUID); err != nil {
		return err
	}
	if err := w.VarInt(s.Type); err != nil {
		return err
	}
	if err := w.Double(s.X); err != nil {
		return err
	}
	if err := w.Double(s.Y); err != nil {
		return err
	}
	if err := w.Double(s.Z); err != nil {
		return err
	}
	if err := w.Ubyte(s.Pitch); err != nil {
		return err
	}
	if err := w.Ubyte(s.Yaw); err != nil {
		return err
	}
	if err := w.Ubyte(s.HeadYaw); err != nil {
		return err
	}
	if err := w.VarInt(s.Data); err != nil {
		return err
	}
	if err := w.Short(s.VelX); err != nil {
		return err
	}
	if err := w.Short(s.VelY); err != nil {
		return err
	}
	return w.Short(s.VelZ)
}

func (s *SpawnEntity) Decode(r io.Reader) error {
	if _, err := r.VarInt(&s.EntityId); err != nil {
		return err
	}
	if err := r.UUID(&s.EntityUUID); err != nil {
		return err
	}
	if _, err := r.VarInt(&s.Type); err != nil {
		return err
	}
	if err := r.Double(&s.X); err != nil {
		return err
	}
	if err := r.Double(&s.Y); err != nil {
		return err
	}
	if err := r.Double(&s.Z); err != nil {
		return err
	}
	if err := r.Ubyte(&s.Pitch); err != nil {
		return err
	}
	if err := r.Ubyte(&s.Yaw); err != nil {
		return err
	}
	if err := r.Ubyte(&s.HeadYaw); err != nil {
		return err
	}
	if _, err := r.VarInt(&s.Data); err != nil {
		return err
	}
	if err := r.Short(&s.VelX); err != nil {
		return err
	}
	if err := r.Short(&s.VelY); err != nil {
		return err
	}
	return r.Short(&s.VelZ)
}
