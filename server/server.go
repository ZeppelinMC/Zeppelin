package server

import (
	"github.com/aimjel/minecraft"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server/network"
	p "github.com/dynamitemc/dynamite/server/player"
	"github.com/dynamitemc/dynamite/server/world"
)

type Server struct {
	Config *ServerConfig
	Logger logger.Logger

	listener *minecraft.Listener

	wrld *world.World
}

func (srv *Server) Start() error {
	for {
		conn, err := srv.listener.Accept()
		if err != nil {
			return err
		}
		go srv.handleNewConn(conn)
	}
}

func (srv *Server) handleNewConn(conn *minecraft.Conn) {
	session := network.NewSession(conn)
	player := p.NewPlayer(session)

	player.JoinDimension(0,
		srv.Config.Hardcore,
		byte(p.Gamemode(srv.Config.Gamemode)),
		srv.wrld.DefaultDimension(),
		srv.wrld.Seed(),
		int32(srv.Config.ViewDistance),
		int32(srv.Config.SimulationDistance),
	)
}

//translate gamemode
