package server

import (
	"sync/atomic"
	"time"

	"github.com/dynamitemc/aether/server/session/std"
	_ "github.com/dynamitemc/aether/server/session/std/handler"

	"github.com/dynamitemc/aether/server/world"

	"github.com/dynamitemc/aether/log"
	"github.com/dynamitemc/aether/net"
	"github.com/dynamitemc/aether/server/session"
)

type Server struct {
	cfg      ServerConfig
	listener *net.Listener
	ticker   Ticker

	world *world.World

	broadcast *session.Broadcast

	entityId atomic.Int32
}

func (srv *Server) Start(ts time.Time) {
	srv.ticker.Start()
	log.Infof("Started server ticker (%d TPS)\n", srv.cfg.TPS)
	log.Infof("Done! (%s)\n", time.Since(ts))
	for {
		conn, err := srv.listener.Accept()
		if err != nil {
			log.Errorln("Server error: ", err)
			return
		}
		log.Infof("[%s] Player attempting to connect: %s (%s)\n", conn.RemoteAddr(), conn.Username(), conn.UUID())
		std.NewStandardSession(conn, srv.entityId.Add(1), srv.world, srv.broadcast).Login()
	}
}
