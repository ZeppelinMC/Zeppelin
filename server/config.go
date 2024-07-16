package server

import (
	"fmt"
	net2 "net"

	"github.com/dynamitemc/aether/log"
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
	Net: ServerConfigNet{
		ServerIP:             net2.IPv4(127, 0, 0, 1),
		ServerPort:           25565,
		CompressionThreshold: -1,
		TPS:                  20,
		EncryptionMode:       encryptionOnline,
	},
	MOTD: "Aether Minecraft Server",
}

type ServerConfigNet struct {
	ServerIP             net2.IP
	ServerPort           int
	CompressionThreshold int32
	TPS                  int
	EncryptionMode       string
}

type ServerConfig struct {
	Net                ServerConfigNet
	MOTD               string
	RenderDistance     int32
	SimulationDistance int32
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
			EnforcesSecureChat: true,
		}),

		IP:                   cfg.Net.ServerIP,
		Port:                 cfg.Net.ServerPort,
		CompressionThreshold: cfg.Net.CompressionThreshold,
		Encrypt:              cfg.Net.EncryptionMode == encryptionYes || cfg.Net.EncryptionMode == encryptionOnline,
		Authenticate:         cfg.Net.EncryptionMode == encryptionOnline,
	}
	listener, err := lcfg.New()
	server := &Server{
		listener:  listener,
		cfg:       cfg,
		world:     world.NewWorld("world"),
		Broadcast: session.NewBroadcast(),
	}

	compstr := "compress everything"
	if cfg.Net.CompressionThreshold > 0 {
		compstr = fmt.Sprintf("compress everything over %d bytes", cfg.Net.CompressionThreshold)
	} else if cfg.Net.CompressionThreshold < 0 {
		compstr = "no compression"
	}

	log.Infof("Compression threshold is %d (%s)\n", cfg.Net.CompressionThreshold, compstr)
	server.createTicker()
	return server, err
}
