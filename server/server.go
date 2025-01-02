package server

import (
	"bytes"
	"fmt"
	"github.com/zeppelinmc/zeppelin/properties"
	"github.com/zeppelinmc/zeppelin/protocol/net/packet/configuration"
	"github.com/zeppelinmc/zeppelin/server/player"
	"image/png"
	"os"
	"runtime"
	"slices"
	"time"

	nnet "net"

	"github.com/zeppelinmc/zeppelin/protocol/net/packet/status"
	"github.com/zeppelinmc/zeppelin/protocol/text"
	"github.com/zeppelinmc/zeppelin/server/command"
	"github.com/zeppelinmc/zeppelin/server/player/state"
	"github.com/zeppelinmc/zeppelin/server/tick"

	"github.com/zeppelinmc/zeppelin/server/world"
	_ "github.com/zeppelinmc/zeppelin/server/world/terrain"

	"github.com/zeppelinmc/zeppelin/protocol/net"
	"github.com/zeppelinmc/zeppelin/util/log"
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
		log.Warnln("Server is not enforcing authentication. The server will let anyone log as any username and potentially harm the server. Proceed with caution")
	}

	listener, err := lcfg.New()
	server := &Server{
		listener:   listener,
		cfg:        cfg,
		World:      world,
		Players:    state.NewPlayerEntityManager(),
		stopLoop:   make(chan struct{}),
		playerList: player.NewPlayerList(0),
	}
	server.icon, err = os.ReadFile("server-icon.png")
	if err == nil {
		img, err := png.Decode(bytes.NewReader(server.icon))
		if err != nil {
			log.Warn("Server icon must be a 64x64 image in the PNG format")
		}
		b := img.Bounds()
		if b.Dx() != 64 || b.Dy() != 64 {
			log.Warn("Server icon must be a 64x64 image in the PNG format")
		}
	}

	//server.Console = &Console{Server: server}
	//server.World.Broadcast.AddDummy(server.Console)
	server.listener.SetStatusProvider(server.provideStatus)

	if server.cfg.EnforceSecureProfile && !cfg.OnlineMode {
		server.cfg.EnforceSecureProfile = false
	}

	compstr := "compress everything"
	if cfg.NetworkCompressionThreshold > 0 {
		compstr = fmt.Sprintf("compress everything starting from %d bytes", cfg.NetworkCompressionThreshold)
	} else if cfg.NetworkCompressionThreshold < 0 {
		compstr = "no compression"
	}

	log.Infolnf("Network compression threshold is %d (%s)", cfg.NetworkCompressionThreshold, compstr)
	//server.createTicker()
	return server, err
}

type Server struct {
	cfg         properties.ServerProperties
	listener    *net.Listener
	TickManager *tick.TickManager

	playerList player.PlayerList

	timeStart time.Time

	World *world.World

	icon []byte

	CommandManager        *command.Manager
	onConnectionIntercept func(conn *net.Conn, stop *bool)

	closed   bool
	stopLoop chan struct{}

	Players *state.PlayerEntityManager
}

func (srv *Server) setOnConnectionIntercept(i func(conn *net.Conn, stop *bool)) {
	srv.onConnectionIntercept = i
}

func (srv *Server) Listener() *net.Listener {
	return srv.listener
}

func (srv *Server) provideStatus(*net.Conn) status.StatusResponseData {
	//count := srv.World.Broadcast.NumSession()
	count := 0
	max := srv.cfg.MaxPlayers
	if max == -1 {
		max = count + 1
	}
	return status.StatusResponseData{
		Version: status.StatusVersion{
			Name:     "Zeppelin 1.21.3",
			Protocol: net.ProtocolVersion,
		},
		Description: text.Unmarshal(srv.cfg.MOTD, srv.cfg.ChatFormatter.Rune()),
		Players: status.StatusPlayers{
			Max:    max,
			Online: count,
			//Sample: srv.World.Broadcast.Sample(),
		},
		Favicon:            srv.icon,
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
	if slices.Index(os.Args, "--no-plugins") == -1 {
		if runtime.GOOS == "darwin" || runtime.GOOS == "linux" || runtime.GOOS == "freebsd" {
			log.Infoln("Loading plugins")
			srv.loadPlugins()
		}
	}

	log.Info("Preparing world spawn... ")
	ow := srv.World.Dimension("minecraft:overworld")
	var chunks int32

	spawnCX, spawnCZ := srv.World.Data.SpawnX>>4, srv.World.Data.SpawnZ>>4
	for x := spawnCX - srv.cfg.ViewDistance; x < spawnCX+srv.cfg.ViewDistance; x++ {
		for z := spawnCZ - srv.cfg.ViewDistance; z < spawnCZ+srv.cfg.ViewDistance; z++ {
			_, err := ow.GetChunk(x, z)
			if err == nil {
				chunks++
			}
		}
	}

	fmt.Printf("cached %d chunks\n\r", chunks)

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
		if srv.onConnectionIntercept != nil {
			var stop bool
			srv.onConnectionIntercept(conn, &stop)
			if stop {
				continue
			}
		}
		srv.handleNewConnection(conn)
	}
}

func (srv *Server) handleNewConnection(conn *net.Conn) {
	log.Infolnf("%sPlayer attempting to connect: %s (%s)", log.FormatAddr(srv.cfg.LogIPs, conn.RemoteAddr()), conn.Username(), conn.UUID())
	if p := srv.playerList.PlayerByUsername(conn.Username()); p != nil {
		_ = conn.WritePacket(&configuration.Disconnect{
			Reason: text.TextComponent{Text: "You are already connected to the server from another session. Please disconnect then try again"},
		})
		return
	}
	playerData, err := srv.World.PlayerData(conn.UUID().String())
	if err != nil {
		playerData = srv.World.NewPlayerData(conn.UUID())
	}

	playerState := srv.Players.New(playerData)

	go srv.playerList.New(conn, playerState, &srv.World.DimensionManager, &srv.World.Level, &srv.cfg, srv.CommandManager).Login()
}

func (srv *Server) Stop() {
	log.InfolnClean("Stopping server")
	srv.closed = true
	srv.listener.Close()
	//srv.World.Broadcast.DisconnectAll(text.TextComponent{Text: "Server closed"})

	log.InfolnClean("Saving player data")
	srv.Players.SaveAll()

	srv.World.Save()
	log.InfolnfClean("Server lasted for %s", srv.formatTimestart())
	srv.stopLoop <- struct{}{}
}

func (srv *Server) formatTimestart() string {
	sub := time.Since(srv.timeStart)
	return fmt.Sprintf("%02dhrs %02dmins, %02dsecs", int(sub.Hours())%60, int(sub.Minutes())%60, int(sub.Seconds())%60)
}

func (srv *Server) createTicker() {
	//srv.TickManager = tick.New(20, srv.World.Broadcast)
	srv.TickManager.AddNew(func() {
		srv.World.IncrementTime()
	})
}
