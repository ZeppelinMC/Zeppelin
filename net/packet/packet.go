package packet

import (
	"github.com/zeppelinmc/zeppelin/net/io"
)

type Packet interface {
	ID() int32
	Decode(io.Reader) error
	Encode(io.Writer) error
}

type UnknownPacket struct {
	Id      int32
	Length  int32
	Payload io.Reader
}

func (u UnknownPacket) ID() int32 {
	return u.Id
}

func (u UnknownPacket) Decode(io.Reader) error {
	return nil
}

func (u UnknownPacket) Encode(io.Writer) error {
	return nil
}

type EmptyPacket struct {
}

func (pk EmptyPacket) Encode(io.Writer) error {
	return nil
}

func (pk EmptyPacket) Decode(io.Reader) error {
	return nil
}
