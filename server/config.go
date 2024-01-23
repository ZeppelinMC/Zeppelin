package server

import (
	"github.com/pelletier/go-toml/v2"
	"os"
)

func LoadConfig(name string) (*Config, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	return &cfg, toml.NewDecoder(f).Decode(&cfg)
}

type Config struct {
	Address              string
	OnlineMode           bool
	CompressionThreshold int32

	MaxPlayers        int
	Description       string
	EnforceSecureChat bool
	PreviewsChat      bool

	ViewDistance int32
}

var DefaultConfig = Config{
	Address: "localhost:25565",

	ViewDistance: 12,
}
