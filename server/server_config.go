package server

import (
	"os"
	"sync"

	"github.com/pelletier/go-toml/v2"

	"github.com/aimjel/minecraft"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/player"
	"github.com/dynamitemc/dynamite/server/world"
)

func LoadConfig(path string, data *Config) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	return toml.NewDecoder(file).Decode(data)
}

func Listen(cfg *Config, address string, logger *logger.Logger, commandGraph *commands.Graph) (*Server, error) {
	lnCfg := minecraft.ListenConfig{
		Status: minecraft.NewStatus(minecraft.Version{
			Text:     "DynamiteMC 1.20.1",
			Protocol: 763,
		}, cfg.MaxPlayers, cfg.MOTD),
		OnlineMode:           cfg.Online,
		CompressionThreshold: int32(cfg.CompressionThreshold),
		Messages: &minecraft.Messages{
			OnlineMode:     cfg.Messages.OnlineMode,
			ProtocolTooNew: cfg.Messages.ProtocolNew,
			ProtocolTooOld: cfg.Messages.ProtocolOld,
		},
	}
	//web.SetMaxPlayers(cfg.MaxPlayers)

	ln, err := lnCfg.Listen(address)
	if err != nil {
		return nil, err
	}
	w, err := world.OpenWorld("world", cfg.Superflat)
	if err != nil {
		world.CreateWorld(cfg.Hardcore)
		logger.Error("Failed to load world: %s", err)
		os.Exit(1)
	}
	w.Gamemode = byte(player.Gamemode(cfg.Gamemode))
	srv := &Server{
		Config:       cfg,
		listener:     ln,
		Logger:       logger,
		World:        w,
		mu:           &sync.RWMutex{},
		Players:      make(map[string]*PlayerController),
		Entities:     make(map[int32]*Entity),
		commandGraph: commandGraph,
		Plugins:      make(map[string]*Plugin),
	}

	logger.Info("Loading player info")
	srv.loadFiles()

	logger.Info("Loading plugins")
	srv.LoadPlugins()

	return srv, nil
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

type Config struct {
	ServerIP             string    `toml:"server_ip"`
	ServerPort           int       `toml:"server_port"`
	ViewDistance         int       `toml:"view_distance"`
	SimulationDistance   int       `toml:"simulation_distance"`
	Superflat            bool      `toml:"superflat"`
	MOTD                 string    `toml:"motd"`
	Whitelist            Whitelist `toml:"whitelist"`
	Web                  Web       `toml:"web"`
	Gamemode             string    `toml:"gamemode"`
	Hardcore             bool      `toml:"hardcore"`
	MaxPlayers           int       `toml:"max_players"`
	Online               bool      `toml:"online_mode"`
	CompressionThreshold int       `toml:"compression_threshold"`
	Tablist              Tablist   `toml:"tablist"`
	Chat                 Chat      `toml:"chat"`
	Messages             Messages  `toml:"messages"`
}
