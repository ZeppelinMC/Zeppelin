package config

import (
	net2 "net"
)

const (
	EncryptionNo     = "disabled"
	EncryptionYes    = "enabled"
	EncryptionOnline = "online"
)

type formatter string

func (f formatter) Rune() rune {
	return rune(f[0])
}

var DefaultConfig = ServerConfig{
	Net: ServerConfigNet{
		ServerIP:             net2.IPv4(127, 0, 0, 1),
		ServerPort:           25565,
		CompressionThreshold: -1,
		TPS:                  20,
		EncryptionMode:       EncryptionOnline,
	},
	Chat: ServerConfigChat{
		ChatMode:   "secure",
		ChatFormat: "<%player%> %message%",
		Formatter:  "&",
	},
	MOTD:               "Zeppelin Minecraft Server",
	RenderDistance:     16,
	SimulationDistance: 16,
	Brand:              "Zeppelin",
}

type ServerConfigNet struct {
	ServerIP             net2.IP `comment:"The ip to listen to connections on"`
	ServerPort           int     `comment:"The port to listen to connections on"`
	CompressionThreshold int32   `comment:"The minimum packet size to compress"`
	TPS                  int     `comment:"Ticks per second"`
	EncryptionMode       string  `comment:"Can be enabled, online or disabled\n Online will also authenticate players when connecting"`
}

type ServerConfig struct {
	Net                ServerConfigNet  `comment:"Network configuration"`
	MOTD               string           `comment:"Message of the day"`
	Chat               ServerConfigChat `comment:"Chat configuration"`
	RenderDistance     int32            `comment:"The radius of chunks to render at a time"`
	SimulationDistance int32            `comment:"The radius of chunks to process at a time"`
	Brand              string           `toml:",commented" comment:"Custom brand name for this server, shown in the debug screen"`
}

type ServerConfigChat struct {
	ChatMode   string    `comment:"Can be secure, disguised, system, or off\n Secure will encrypt all chat messages and enable player reporting, and requires encryption mode set to online\n Disguised will use the default chat format unauthenticated\n System will use the format specified, unauthenticated\n Off will disable the chat entirely"`
	ChatFormat string    `comment:"Can only be used with system chat mode"`
	Formatter  formatter `comment:"Character used for text formatting"`
	Colors     bool      `comment:"Whether to allow color codes or not"`
}
