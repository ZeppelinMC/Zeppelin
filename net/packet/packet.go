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

// packet that writes the specified payload without length prefix
type raw struct {
	payload []byte
	id      int32
}

func (pk raw) ID() int32 {
	return pk.id
}

func (pk raw) Encode(w io.Writer) error {
	return w.FixedByteArray(pk.payload)
}

func (raw) Decode(r io.Reader) error {
	return nil
}

func Raw(id int32, payload []byte) Packet {
	return raw{id: id, payload: payload}
}
