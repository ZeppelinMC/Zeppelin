package net

import (
	"crypto/rsa"
	"net"

	"github.com/dynamitemc/aether/net/io"
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

	privKey *rsa.PrivateKey
}

func (l *Listener) Accept() (*Conn, error) {
	c, err := l.Listener.Accept()
	conn := &Conn{
		Conn:     c,
		listener: l,

		//rd: bufio.NewReaderSize(c, 4096),
	}
	conn.writer = io.NewWriter(conn)

	if err != nil {
		return conn, err
	}

	if !conn.handleHandshake() {
		conn = nil
	}

	return conn, nil
}
