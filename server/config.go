package server

import (
	"aether/net"
	"aether/net/packet/status"
	net2 "net"
)

type ServerConfig struct {
	IP                   net2.IP
	Port                 int
	CompressionThreshold int32
	TPS                  int
}

func (cfg ServerConfig) New() (*Server, error) {
	lcfg := net.Config{
		Status: net.Status(status.StatusResponseData{
			Version: status.StatusVersion{
				Name:     "1.21",
				Protocol: net.ProtocolVersion,
			},
			Description: status.StatusDescription{Text: "welcome to our minecraft server!"},
			Players: status.StatusPlayers{
				Max: 20,
			},
		}),

		IP:                   cfg.IP,
		Port:                 cfg.Port,
		CompressionThreshold: cfg.CompressionThreshold,
	}
	listener, err := lcfg.New()
	server := &Server{
		listener: listener,
		cfg:      cfg,
	}
	server.createTicker()
	return server, err
}
