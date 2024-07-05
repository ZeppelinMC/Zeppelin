package net

import (
	"aether/net/io"
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

	if !conn.handleHandshake() {
		conn = nil
	}

	return conn, nil
}
