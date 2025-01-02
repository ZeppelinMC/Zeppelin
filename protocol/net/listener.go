// Package net provides tools for creating minecraft servers, such as packet sending, registry data, encryption, authentication, compression and more
package net

import (
	"crypto/rsa"
	"fmt"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet"
	"net"

	"github.com/zeppelinmc/zeppelin/protocol/text"
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

	ApprovePlayer func(*Conn) (ok bool, disconnectionReason *text.TextComponent)

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

		packetWriteChan: make(chan packet.Encodeable, l.cfg.PacketWriteChanBuffer),
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
