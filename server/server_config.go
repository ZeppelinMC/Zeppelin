package server

import (
	"errors"
	"math"
	"os"
	"sync"

	"github.com/aimjel/minecraft"
	"github.com/dynamitemc/dynamite/config"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/player"
	"github.com/dynamitemc/dynamite/server/world"
)

func Listen(cfg *config.ServerConfig, address string, logger logger.Logger, commandGraph *commands.Graph) (*Server, error) {
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
		world:        w,
		mu:           &sync.RWMutex{},
		Players:      make(map[string]*PlayerController),
		CommandGraph: commandGraph,
		Plugins:      make(map[string]*Plugin),
	}

	logger.Info("Loading player info")
	var files = []string{"whitelist.json", "banned_players.json", "ops.json", "banned_ips.json"}
	var addresses = []*[]user{&srv.WhitelistedPlayers, &srv.BannedPlayers, &srv.Operators, &srv.BannedIPs}
	for i, file := range files {
		u, err := loadUsers(file)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return nil, err
		}

		*addresses[i] = u
	}

	if cfg.CacheSpawnChunks {
		count := 4 * math.Pow(float64(cfg.ViewDistance), 2)
		logger.Info("Caching spawn chunks (%v)", count)
		s := w.LoadSpawnChunks(int32(cfg.ViewDistance))
		logger.Debug("Cached spawn chunks (%d/%v)", s, count)
	}

	logger.Info("Loading plugins")
	srv.LoadPlugins()

	return srv, nil
}
