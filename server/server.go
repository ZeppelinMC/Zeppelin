package server

import (
	"sync"

	"github.com/aimjel/minecraft"
	"github.com/dynamitemc/dynamite/gui"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/network"
	p "github.com/dynamitemc/dynamite/server/player"
	"github.com/dynamitemc/dynamite/server/world"
	"github.com/dynamitemc/dynamite/util"
)

type Server struct {
	*sync.Mutex
	Config       *ServerConfig
	Logger       logger.Logger
	CommandGraph commands.Graph
	Players      map[string]*p.Player

	listener *minecraft.Listener

	world *world.World
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
	srv.addPlayer(player)

	player.JoinDimension(0,
		srv.Config.Hardcore,
		byte(p.Gamemode(srv.Config.Gamemode)),
		srv.world.DefaultDimension(),
		srv.world.Seed(),
		int32(srv.Config.ViewDistance),
		int32(srv.Config.SimulationDistance),
	)

	if err := session.HandlePackets(); err != nil {
		u := session.Conn.Info.UUID
		uuid := util.ParseUUID(u)

		srv.Logger.Info("[%s] Player %s (%s) has left the server", conn.RemoteAddr().String(), conn.Info.Name, uuid)
		srv.PlayerlistRemove(u)
		gui.RemovePlayer(uuid)
	}
}

func (srv *Server) addPlayer(p *p.Player) {
	uuid := util.ParseUUID(p.Session.Conn.Info.UUID)
	srv.Lock()
	srv.Players[uuid] = p
	srv.Unlock()
	srv.PlayerlistUpdate()
	gui.AddPlayer(p.Session.Conn.Info.Name, uuid)

	srv.Logger.Info("[%s] Player %s (%s) has joined the server", p.Session.Conn.RemoteAddr().String(), p.Session.Conn.Info.Name, uuid)

}

//translate gamemode
