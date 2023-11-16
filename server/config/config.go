package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

func LoadConfig(path string, data *Config) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	return toml.NewDecoder(file).Decode(data)
}

var DefaultConfig = Config{
	ServerIP:             "0.0.0.0",
	ServerPort:           25565,
	ViewDistance:         10,
	SimulationDistance:   10,
	MOTD:                 "A DynamiteMC Minecraft server!",
	Online:               true,
	CompressionThreshold: 256,
	Gamemode:             "survival",
	Hardcore:             false,
	MaxPlayers:           20,
	TPS:                  20,
	Web: Web{
		ServerIP:   "0.0.0.0",
		ServerPort: 8080,
		Password:   "ChangeMe",
		Enable:     false,
	},
	Chat: Chat{
		Colors: false,
		Format: "<%player_prefix%%player%> %message%",
		Enable: true,
	},
}

type Tablist struct {
	Header []string `toml:"header"`
	Footer []string `toml:"footer"`
}

type Web struct {
	ServerIP   string `toml:"server_ip"`
	ServerPort int    `toml:"server_port"`
	Password   string `toml:"password"`
	Enable     bool   `toml:"enable"`
}

type Chat struct {
	Format string `toml:"format"`
	Secure bool   `toml:"secure"`
	Colors bool   `toml:"colors"`
	Enable bool   `toml:"enable"`
}

type Whitelist struct {
	Enforce bool `toml:"enforce"`
	Enable  bool `toml:"enable"`
}

type ResourcePack struct {
	URL    string `toml:"url"`
	Hash   string `toml:"hash"`
	Force  bool   `toml:"force"`
	Enable bool   `toml:"enable"`
}

type Config struct {
	ServerIP             string       `toml:"server_ip"`
	ServerPort           int          `toml:"server_port"`
	ViewDistance         int          `toml:"view_distance"`
	SimulationDistance   int          `toml:"simulation_distance"`
	Superflat            bool         `toml:"superflat"`
	MOTD                 string       `toml:"motd"`
	Gamemode             string       `toml:"gamemode"`
	Hardcore             bool         `toml:"hardcore"`
	MaxPlayers           int          `toml:"max_players"`
	TPS                  int64        `toml:"tps"`
	Online               bool         `toml:"online_mode"`
	CompressionThreshold int          `toml:"compression_threshold"`
	Whitelist            Whitelist    `toml:"whitelist"`
	Web                  Web          `toml:"web"`
	Tablist              Tablist      `toml:"tablist"`
	Chat                 Chat         `toml:"chat"`
	ResourcePack         ResourcePack `toml:"resource_pack"`
}
