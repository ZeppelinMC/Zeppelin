package server

import (
	net2 "net"

	"github.com/dynamitemc/aether/net"
	"github.com/dynamitemc/aether/net/packet/status"
	"github.com/dynamitemc/aether/server/session"
	"github.com/dynamitemc/aether/server/world"
)

const (
	encryptionNo     = "disabled"
	encryptionYes    = "enabled"
	encryptionOnline = "online"
)

var DefaultConfig = ServerConfig{
	ServerIP:             net2.IPv4(127, 0, 0, 1),
	ServerPort:           25565,
	CompressionThreshold: -1,
	TPS:                  20,
	EncryptionMode:       encryptionOnline,
	MOTD:                 "Aether Minecraft Server",
}

type ServerConfig struct {
	ServerIP             net2.IP
	ServerPort           int
	CompressionThreshold int32
	TPS                  int
	EncryptionMode       string
	MOTD                 string
}

func (cfg ServerConfig) New() (*Server, error) {
	lcfg := net.Config{
		Status: net.Status(status.StatusResponseData{
			Version: status.StatusVersion{
				Name:     "1.21",
				Protocol: net.ProtocolVersion,
			},
			Description: status.StatusDescription{Text: cfg.MOTD},
			Players: status.StatusPlayers{
				Max: 20,
			},
		}),

		IP:                   cfg.ServerIP,
		Port:                 cfg.ServerPort,
		CompressionThreshold: cfg.CompressionThreshold,
		Encrypt:              cfg.EncryptionMode == encryptionYes || cfg.EncryptionMode == encryptionOnline,
		Authenticate:         cfg.EncryptionMode == encryptionOnline,
	}
	listener, err := lcfg.New()
	server := &Server{
		listener:  listener,
		cfg:       cfg,
		world:     world.NewWorld("world"),
		broadcast: session.NewBroadcast(),
	}
	server.createTicker()
	return server, err
}
