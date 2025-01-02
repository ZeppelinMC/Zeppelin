package packet

import "github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"

// Encodeable is a clientbound packet
type Encodeable interface {
	ID() int32
	Encode(encoding.Writer) error
}

// Decodeable is a serverbound packet
type Decodeable interface {
	ID() int32
	Decode(encoding.Reader) error
}

// Error is a dummy decodeable packet that is simply an error
type Error struct {
	Error error
}

func (Error) ID() int32                    { return -1 }
func (Error) Decode(encoding.Reader) error { return nil }

var _ Decodeable = (*Error)(nil)

type UnknownPacket struct {
	Id      int32
	Length  int32
	Payload encoding.Reader
}

func (u UnknownPacket) ID() int32 {
	return u.Id
}

func (u UnknownPacket) Decode(encoding.Reader) error {
	return nil
}

func (u UnknownPacket) Encode(encoding.Writer) error {
	return nil
}

type EmptyPacket struct {
}

func (pk EmptyPacket) Encode(encoding.Writer) error {
	return nil
}

func (pk EmptyPacket) Decode(encoding.Reader) error {
	return nil
}
