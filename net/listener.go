package net

import (
	"aether/net/io"
	"aether/net/packet/handshake"
	"aether/net/packet/status"
	"fmt"
	"net"
)

const (
	ProtocolVersion = 767
)

const (
	HandshakingState = iota
	StatusState
	LoginState
	ConfigurationState
	PlayState
)

type Listener struct {
	net.Listener
	Config
}

func (l *Listener) Accept() (*Conn, error) {
	c, err := l.Listener.Accept()
	conn := &Conn{
		Conn:   c,
		reader: io.NewReader(c),
		writer: io.NewWriter(c),

		listener: l,
	}
	if err != nil {
		return conn, err
	}
	pk, err := conn.ReadPacket()
	if err != nil {
		return conn, err
	}
	handshaking, ok := pk.(*handshake.Handshaking)
	if !ok {
		return conn, fmt.Errorf("expected packet Handshaking, got %T", pk)
	}

	switch handshaking.NextState {
	case handshake.Status:
		conn.state = StatusState
		pk, err := conn.ReadPacket()
		if err != nil {
			return conn, err
		}
		_, ok := pk.(*status.StatusRequest)
		if !ok {
			return conn, fmt.Errorf("expected packet StatusRequest, got %T", pk)
		}
		if err := conn.WritePacket(&status.StatusResponse{Data: l.Status()}); err != nil {
			return conn, err
		}

		pk, err = conn.ReadPacket()
		if err != nil {
			return conn, err
		}
		p, ok := pk.(*status.Ping)
		if !ok {
			return conn, fmt.Errorf("expected packet PingRequest, got %T", pk)
		}
		if err := conn.WritePacket(p); err != nil {
			return conn, err
		}
	}

	return conn, nil
}
