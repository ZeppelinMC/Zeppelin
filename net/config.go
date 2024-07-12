package net

import (
	"crypto/rand"
	"crypto/rsa"
	"net"

	"github.com/dynamitemc/aether/net/packet/status"
)

type Config struct {
	Status func() status.StatusResponseData

	IP                   net.IP
	Port                 int
	CompressionThreshold int32
	Encrypt              bool
	Authenticate         bool
}

func (c Config) New() (*Listener, error) {
	l, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP: c.IP, Port: c.Port,
	})
	if err != nil {
		return nil, err
	}
	lis := &Listener{
		Config:   c,
		Listener: l,
	}
	lis.privKey, err = rsa.GenerateKey(rand.Reader, 1024)

	return lis, err
}

func Status(s status.StatusResponseData) func() status.StatusResponseData {
	return func() status.StatusResponseData {
		return s
	}
}
