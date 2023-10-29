package server

import (
	"os"

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
	}

	if cfg.Chat.Secure && !cfg.Online {
		logger.Warn("Secure chat doesn't work on offline mode")
		cfg.Chat.Secure = false
	}
	if cfg.Chat.Secure && cfg.Chat.Format != "" {
		logger.Warn("Secure chat overrides the chat format")
	}
	if cfg.TPS < 20 {
		logger.Warn("TPS must be at least 20")
		cfg.TPS = 20
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
		players:      make(map[string]*Session),
		entities:     make(map[int32]*Entity),
		commandGraph: commandGraph,
	}

	logger.Info("Loading player info")
	srv.loadFiles()

	go srv.tickLoop()

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
	Whitelist            Whitelist    `toml:"whitelist"`
	Web                  Web          `toml:"web"`
	Gamemode             string       `toml:"gamemode"`
	Hardcore             bool         `toml:"hardcore"`
	MaxPlayers           int          `toml:"max_players"`
	TPS                  int64        `toml:"tps"`
	Online               bool         `toml:"online_mode"`
	CompressionThreshold int          `toml:"compression_threshold"`
	Tablist              Tablist      `toml:"tablist"`
	Chat                 Chat         `toml:"chat"`
	ResourcePack         ResourcePack `toml:"resource_pack"`
}
