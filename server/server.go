package server

import (
	"runtime"
	"sync/atomic"
	"time"

	"github.com/dynamitemc/aether/server/player"
	"github.com/dynamitemc/aether/server/session/std"
	_ "github.com/dynamitemc/aether/server/session/std/handler"
	"github.com/dynamitemc/aether/text"

	"github.com/dynamitemc/aether/server/world"

	"github.com/dynamitemc/aether/log"
	"github.com/dynamitemc/aether/net"
	"github.com/dynamitemc/aether/server/session"
)

type Server struct {
	cfg      ServerConfig
	listener *net.Listener
	ticker   Ticker

	World *world.World

	Broadcast *session.Broadcast

	entityId atomic.Int32

	closed bool
}

func (srv *Server) Config() ServerConfig {
	return srv.cfg
}

func (srv *Server) NewEntityId() int32 {
	return srv.entityId.Add(1)
}

func (srv *Server) Start(ts time.Time) {
	if runtime.GOOS != "darwin" && runtime.GOOS != "linux" && runtime.GOOS != "freebsd" {
		log.Warnf("Your platform doesn't support plugins (yet) bozo! Supported platforms: Linux, macOS, and FreeBSD. You are using %s\n> ", runtime.GOOS)
	} else {
		log.Infoln("Loading plugins")
		srv.loadPlugins()
	}
	srv.ticker.Start()
	log.Infolnf("Started server ticker (%d TPS)", srv.cfg.Net.TPS)
	log.Infolnf("Done! (%s)", time.Since(ts))
	for {
		conn, err := srv.listener.Accept()
		if err != nil {
			if !srv.closed {
				log.Errorln("Server error: ", err)
			}
			return
		}
		log.Infolnf("[%s] Player attempting to connect: %s (%s)", conn.RemoteAddr(), conn.Username(), conn.UUID())
		player := player.NewPlayer(srv.entityId.Add(1))
		std.NewStandardSession(conn, player, srv.World, srv.Broadcast).Login()
	}
}

func (srv *Server) Stop() {
	log.InfolnClean("Stopping server")
	srv.closed = true
	srv.Broadcast.DisconnectAll(text.TextComponent{Text: "Server closed"})
	srv.listener.Close()
}
