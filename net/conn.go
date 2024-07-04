package net

import (
	"aether/net/io"
	"aether/net/packet"
	"bytes"
	"net"
)

type Conn struct {
	net.Conn

	listener *Listener

	reader io.Reader
	writer io.Writer

	state int
}

func (conn *Conn) WritePacket(pk packet.Packet) error {
	if conn.listener.CompressionThreshold < 0 {
		var buf = new(bytes.Buffer)
		w := io.NewWriter(buf)
		if err := w.VarInt(pk.ID()); err != nil {
			return err
		}
		if err := pk.Encode(w); err != nil {
			return err
		}
		if err := conn.writer.VarInt(int32(buf.Len())); err != nil {
			return err
		}
		return conn.writer.FixedByteArray(buf.Bytes())
	} else {
		//compression
	}
	return nil
}

func (conn *Conn) ReadPacket() (packet.Packet, error) {
	rd := conn.reader
	var (
		length, packetId int32
		data             []byte
	)
	if conn.listener.CompressionThreshold < 0 {
		if _, err := rd.VarInt(&length); err != nil {
			return nil, err
		}
		vii, err := rd.VarInt(&packetId)
		if err != nil {
			return nil, err
		}
		length -= int32(vii)
		data = make([]byte, length)
		if err := rd.FixedByteArray(data); err != nil {
			return nil, err
		}

		rd = io.NewReader(bytes.NewReader(data))
	} else {
		//compression
	}

	var pk packet.Packet
	pc, ok := serverboundPool[conn.state][packetId]
	if !ok {
		return packet.UnknownPacket{
			Id:      packetId,
			Length:  length,
			Payload: rd,
		}, nil
	} else {
		pk = pc()
		err := pk.Decode(rd)
		return pk, err
	}
}
