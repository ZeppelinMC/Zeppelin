package net

import (
	"crypto/rsa"
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
	cfg Config

	newConns chan *Conn
	err      chan error

	privKey *rsa.PrivateKey
}

func (l *Listener) SetStatusProvider(p StatusProvider) {
	l.cfg.Status = p
}

func (l *Listener) StatusProvider() StatusProvider {
	return l.cfg.Status
}

func (l *Listener) listen() {
	for {
		c, err := l.Listener.Accept()
		if err != nil {
			l.err <- err
			close(l.newConns)
			return
		}
		conn := l.newConn(c)

		go func() {
			if conn.handleHandshake() {
				l.newConns <- conn
			}
		}()
	}
}

func (l *Listener) newConn(c net.Conn) *Conn {
	conn := &Conn{
		Conn:     c,
		listener: l,
	}

	return conn
}

func (l *Listener) Close() error {
	close(l.newConns)
	l.err <- fmt.Errorf("listener closed")

	return nil
}

func (l *Listener) Accept() (*Conn, error) {
	conn, ok := <-l.newConns
	if !ok {
		return nil, <-l.err
	}

	return conn, nil
}
