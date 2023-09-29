package config

import (
	"errors"
	"io/fs"
	"os"

	"github.com/pelletier/go-toml/v2"
)

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

type Messages struct {
	NotInWhitelist          string `toml:"not_in_whitelist"`
	Banned                  string `toml:"banned"`
	ServerFull              string `toml:"server_full"`
	AlreadyPlaying          string `toml:"already_playing"`
	PlayerJoin              string `toml:"player_join"`
	PlayerLeave             string `toml:"player_leave"`
	UnknownCommand          string `toml:"unknown_command"`
	ProtocolNew             string `toml:"protocol_new"`
	ProtocolOld             string `toml:"protocol_old"`
	InsufficientPermissions string `toml:"insufficient_permissions"`
	ReloadComplete          string `toml:"reload_complete"`
	ServerClosed            string `toml:"server_closed"`
	OnlineMode              string `toml:"online_mode"`
}

type Chat struct {
	Format string `toml:"format"`
	Colors bool   `toml:"colors"`
	Enable bool   `toml:"enable"`
}

type Whitelist struct {
	Enforce bool `toml:"enforce"`
	Enable  bool `toml:"enable"`
}

type ServerConfig struct {
	ServerIP           string    `toml:"server_ip"`
	ServerPort         int       `toml:"server_port"`
	ViewDistance       int       `toml:"view_distance"`
	SimulationDistance int       `toml:"simulation_distance"`
	MOTD               string    `toml:"motd"`
	Whitelist          Whitelist `toml:"whitelist"`
	Web                Web       `toml:"web"`
	Gamemode           string    `toml:"gamemode"`
	Hardcore           bool      `toml:"hardcore"`
	MaxPlayers         int       `toml:"max_players"`
	Online             bool      `toml:"online_mode"`
	Tablist            Tablist   `toml:"tablist"`
	Chat               Chat      `toml:"chat"`
	Messages           Messages  `toml:"messages"`
}

func LoadConfig(path string, data *ServerConfig) {
	file, err := os.Open(path)
	if errors.Is(err, fs.ErrNotExist) {
		config := defaultConfig()
		CreateConfig(path, config)
	} else {
		defer file.Close()
		decoder := toml.NewDecoder(file)
		decoder.Decode(data)
	}
}

func CreateConfig(path string, config ServerConfig) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := toml.NewEncoder(file)
	return encoder.Encode(config)
}

func defaultConfig() ServerConfig {
	return ServerConfig{
		ServerIP:           "0.0.0.0",
		ServerPort:         25565,
		ViewDistance:       10,
		SimulationDistance: 10,
		MOTD:               "A DynamiteMC Minecraft server!",
		Online:             true,
		Whitelist: Whitelist{
			Enforce: false,
			Enable:  false,
		},
		Gamemode:   "survival",
		Hardcore:   false,
		MaxPlayers: 20,
		Messages: Messages{
			NotInWhitelist:          "You are not whitelisted.",
			Banned:                  "You are banned from this server.",
			ServerFull:              "The server is full.",
			AlreadyPlaying:          "You are already playing on this server with a different client.",
			PlayerJoin:              "§e%player% has joined the game",
			PlayerLeave:             "§e%player% has left the game",
			UnknownCommand:          "§cUnknown command. Please use '/help' for a list of commands.",
			ProtocolNew:             "Your protocol is too new!",
			ProtocolOld:             "Your protocol is too old!",
			InsufficientPermissions: "§cYou aren't permitted to use this command.",
			ReloadComplete:          "§aReload complete.",
			ServerClosed:            "Server closed.",
			OnlineMode:              "The server is in online mode.",
		},
		Web: Web{
			ServerIP:   "0.0.0.0",
			ServerPort: 8080,
			Password:   "ChangeMe",
			Enable:     false,
		},
		Tablist: Tablist{
			Header: []string{},
			Footer: []string{},
		},
		Chat: Chat{
			Colors: false,
			Format: "<%player_prefix%%player%> %message%",
			Enable: true,
		},
	}
}
