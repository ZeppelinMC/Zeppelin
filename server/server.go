package server

import (
	"aether/log"
	"aether/net"
	"aether/server/session"
	"aether/server/world"
	"time"
)

type Server struct {
	cfg      ServerConfig
	listener *net.Listener
	ticker   Ticker

	world *world.World
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
		if conn == nil {
			continue
		}
		log.Infof("[%s] Player attempting to connect: %s (%s)\n", conn.RemoteAddr(), conn.Username(), conn.UUID())
		session.NewSession(conn, 1, srv.world).Login()
	}
}
