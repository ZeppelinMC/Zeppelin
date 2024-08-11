package server

import (
	"fmt"
	"runtime"
	"time"

	nnet "net"

	"github.com/zeppelinmc/zeppelin/net/packet/configuration"
	"github.com/zeppelinmc/zeppelin/net/packet/status"
	"github.com/zeppelinmc/zeppelin/properties"
	"github.com/zeppelinmc/zeppelin/server/command"
	"github.com/zeppelinmc/zeppelin/server/player"
	"github.com/zeppelinmc/zeppelin/server/session/std"
	_ "github.com/zeppelinmc/zeppelin/server/session/std/handler"
	"github.com/zeppelinmc/zeppelin/server/tick"
	"github.com/zeppelinmc/zeppelin/text"
	"github.com/zeppelinmc/zeppelin/util"

	"github.com/zeppelinmc/zeppelin/server/world"
	_ "github.com/zeppelinmc/zeppelin/server/world/terrain"

	"github.com/zeppelinmc/zeppelin/log"
	"github.com/zeppelinmc/zeppelin/net"
)

// Creates a new server instance using the specified config, returns an error if unable to bind listener
func New(cfg properties.ServerProperties, world *world.World) (*Server, error) {
	var ip = nnet.ParseIP(cfg.ServerIp)
	if ip == nil {
		ip = nnet.IPv4(0, 0, 0, 0)
	}

	lcfg := net.Config{
		IP:                   ip,
		Port:                 int(cfg.ServerPort),
		CompressionThreshold: int32(cfg.NetworkCompressionThreshold),
		Encrypt:              cfg.EnableEncryption || cfg.OnlineMode,
		Authenticate:         cfg.OnlineMode,
		AcceptTransfers:      cfg.AcceptTransfers,
	}
	if !cfg.OnlineMode {
		log.Warnln("Server is running in offline mode. The server will let anyone log as any username and potentially harm the server. Proceed with caution")
	}

	listener, err := lcfg.New()
	server := &Server{
		listener: listener,
		cfg:      cfg,
		World:    world,
		Players:  player.NewPlayerManager(),
		stopLoop: make(chan struct{}),
	}
	server.Console = &Console{Server: server}
	server.World.Broadcast.AddDummy(server.Console)
	server.listener.SetStatusProvider(server.provideStatus)

	if server.cfg.EnforceSecureProfile && !server.cfg.OnlineMode {
		server.cfg.EnforceSecureProfile = false
	}

	compstr := "compress everything"
	if cfg.NetworkCompressionThreshold > 0 {
		compstr = fmt.Sprintf("compress everything starting from %d bytes", cfg.NetworkCompressionThreshold)
	} else if cfg.NetworkCompressionThreshold < 0 {
		compstr = "no compression"
	}

	log.Infof("Compression threshold is %d (%s)\n", cfg.NetworkCompressionThreshold, compstr)
	server.createTicker()
	return server, err
}

type Server struct {
	cfg         properties.ServerProperties
	listener    *net.Listener
	TickManager *tick.TickManager

	timeStart time.Time

	World *world.World

	Console *Console

	CommandManager *command.Manager

	closed   bool
	stopLoop chan struct{}

	Players *player.PlayerManager
}

func (srv *Server) provideStatus() status.StatusResponseData {
	count := srv.World.Broadcast.NumSession()
	max := srv.cfg.MaxPlayers
	if max == -1 {
		max = count + 1
	}
	return status.StatusResponseData{
		Version: status.StatusVersion{
			Name:     "Zeppelin 1.21",
			Protocol: net.ProtocolVersion,
		},
		Description: text.Unmarshal(srv.cfg.MOTD, srv.cfg.ChatFormatter.Rune()),
		Players: status.StatusPlayers{
			Max:    max,
			Online: count,
			Sample: srv.World.Broadcast.Sample(),
		},
		EnforcesSecureChat: srv.cfg.EnforceSecureProfile,
	}
}

func (srv *Server) SetStatusProvider(sp net.StatusProvider) {
	srv.listener.SetStatusProvider(sp)
}

func (srv *Server) Properties() properties.ServerProperties {
	return srv.cfg
}

func (srv *Server) Start(ts time.Time) {
	if !util.HasArgument("--no-plugins") {
		if runtime.GOOS == "darwin" || runtime.GOOS == "linux" || runtime.GOOS == "freebsd" {
			log.Infoln("Loading plugins")
			srv.loadPlugins()
		}
	}
	srv.timeStart = ts
	log.Infolnf("Done! (%s)", time.Since(ts))
	for {
		conn, err := srv.listener.Accept()
		if err != nil {
			if !srv.closed {
				log.Errorln("Server error: ", err)
			}
			<-srv.stopLoop
			return
		}
		srv.handleNewConnection(conn)
	}
}

func (srv *Server) handleNewConnection(conn *net.Conn) {
	log.Infolnf("[%s] Player attempting to connect: %s (%s)", conn.RemoteAddr(), conn.Username(), conn.UUID())
	if _, ok := srv.World.Broadcast.SessionByUsername(conn.Username()); ok {
		conn.WritePacket(&configuration.Disconnect{
			Reason: text.TextComponent{Text: "You are already connected to the server from another session. Please disconnect then try again"},
		})
		return
	}
	playerData, err := srv.World.PlayerData(conn.UUID().String())
	if err != nil {
		playerData = srv.World.NewPlayerData(conn.UUID())
	}

	player := srv.Players.New(playerData)
	std.New(conn, player, srv.World, srv.World.Broadcast, srv.cfg, func() net.StatusProvider {
		return srv.listener.StatusProvider()
	}, srv.CommandManager, srv.TickManager).Configure()
}

func (srv *Server) Stop() {
	log.InfolnClean("Stopping server")
	srv.closed = true
	srv.listener.Close()
	srv.World.Broadcast.DisconnectAll(text.TextComponent{Text: "Server closed"})

	log.InfolnClean("Saving player data")
	srv.Players.SaveAll()

	//srv.World.Save()
	log.InfolnfClean("Server lasted for %s", srv.formatTimestart())
	srv.stopLoop <- struct{}{}
}

func (srv *Server) formatTimestart() string {
	sub := time.Since(srv.timeStart)
	return fmt.Sprintf("%02dhrs %02dmins, %02dsecs", int(sub.Hours())%60, int(sub.Minutes())%60, int(sub.Seconds())%60)
}

func (srv *Server) createTicker() {
	srv.TickManager = tick.New(20, srv.World.Broadcast)
}
