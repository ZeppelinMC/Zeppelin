package net

import (
	"net"

	"github.com/dynamitemc/aether/net/packet/status"
)

type Config struct {
	Status func() status.StatusResponseData

	IP                   net.IP
	Port                 int
	CompressionThreshold int32
}

func (c Config) New() (*Listener, error) {
	l, err := net.ListenTCP("tcp", &net.TCPAddr{
		IP: c.IP, Port: c.Port,
	})
	lis := &Listener{
		Config:   c,
		Listener: l,
	}

	return lis, err
}

func Status(s status.StatusResponseData) func() status.StatusResponseData {
	return func() status.StatusResponseData {
		return s
	}
}
