package packet

import "github.com/zeppelinmc/zeppelin/protocol/net/io/encoding"

// A clientbound packet
type Encodeable interface {
	ID() int32
	Encode(encoding.Writer) error
}

// A serverbound packet
type Decodeable interface {
	ID() int32
	Decode(encoding.Reader) error
}

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
