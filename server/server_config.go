package server

import (
	"errors"
	"os"
	"sync"

	"github.com/aimjel/minecraft"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/world"
	"github.com/dynamitemc/dynamite/util"
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

func (cfg *ServerConfig) Listen(address string, logger logger.Logger, commandGraph *commands.Graph) (*Server, error) {
	lnCfg := minecraft.ListenConfig{
		Status: minecraft.NewStatus(minecraft.Version{
			Text:     "DynamiteMC 1.20.1",
			Protocol: 763,
		}, cfg.MaxPlayers, cfg.MOTD),
		OnlineMode:           cfg.Online,
		CompressionThreshold: 256,
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
	w, err := world.OpenWorld(util.GetArg("worldpath", "world"))
	if err != nil {
		logger.Error("Failed to load world: %s", err)
		os.Exit(1)
	}
	srv := &Server{
		Config:       cfg,
		listener:     ln,
		Logger:       logger,
		world:        w,
		mu:           &sync.RWMutex{},
		Players:      make(map[string]*PlayerController),
		CommandGraph: *commandGraph,
		Plugins:      make(map[string]*Plugin),
	}

	var files = []string{"whitelist.json", "banned_players.json", "ops.json", "banned_ips.json"}
	var addresses = []*[]user{&srv.WhitelistedPlayers, &srv.BannedPlayers, &srv.Operators, &srv.BannedIPs}
	for i, file := range files {
		u, err := loadUsers(file)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}

		*addresses[i] = u
	}

	logger.Debug("Loaded player info")
	return srv, nil
}
