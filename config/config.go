package config

import (
	"errors"
	"io/fs"
	"os"

	"github.com/dynamitemc/dynamite/server"
	"github.com/pelletier/go-toml/v2"
)

func LoadConfig(path string, data *server.ServerConfig) {
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

func CreateConfig(path string, config server.ServerConfig) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := toml.NewEncoder(file)
	return encoder.Encode(config)
}

func defaultConfig() server.ServerConfig {
	return server.ServerConfig{
		ServerIP:           "0.0.0.0",
		ServerPort:         25565,
		ViewDistance:       10,
		SimulationDistance: 10,
		MOTD:               "A DynamiteMC Minecraft server!",
		Online:             true,
		Whitelist: server.Whitelist{
			Enforce: false,
			Enable:  false,
		},
		Gamemode:   "survival",
		Hardcore:   false,
		MaxPlayers: 20,
		Messages: server.Messages{
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
		GUI: server.GUI{
			ServerIP:   "0.0.0.0",
			ServerPort: 8080,
			Password:   "ChangeMe",
			Enable:     false,
		},
		Tablist: server.Tablist{
			Header: []string{},
			Footer: []string{},
		},
		Chat: server.Chat{
			Colors: false,
			Format: "<%player_prefix%%player%> %message%",
			Enable: true,
		},
	}
}
