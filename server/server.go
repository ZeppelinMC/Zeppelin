package server

import (
	"encoding/hex"
	"sync"

	"github.com/dynamitemc/dynamite/util"

	"github.com/aimjel/minecraft"
	//"github.com/dynamitemc/dynamite/web"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/player"
	"github.com/dynamitemc/dynamite/server/world"
)

type Server struct {
	Config       *ServerConfig
	Logger       logger.Logger
	CommandGraph commands.Graph

	Plugins map[string]*Plugin

	// Players mapped by UUID
	Players map[string]*PlayerController

	WhitelistedPlayers,
	Operators,
	BannedPlayers,
	BannedIPs []user

	listener *minecraft.Listener

	world *world.World

	mu *sync.RWMutex
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
	if srv.ValidateConn(conn) {
		return
	}

	plyr := player.New()
	sesh := New(conn, plyr)
	cntrl := &PlayerController{player: plyr, session: sesh}
	cntrl.UUID = util.AddDashesToUUID(hex.EncodeToString(conn.Info.UUID[:]))
	if err := cntrl.JoinDimension(srv.world.DefaultDimension()); err != nil {
		//TODO log error
		conn.Close(err)
		panic(err)
		return
	}

	if err := cntrl.SendAvailableCommands(srv.CommandGraph.Data()); err != nil {
		//TODO log error
		conn.Close(err)
		return
	}

	srv.addPlayer(cntrl)
	if err := sesh.HandlePackets(); err != nil {
		u := cntrl.UUID

		srv.Logger.Info("[%s] Player %s (%s) has left the server", conn.RemoteAddr().String(), conn.Info.Name, u)
		srv.PlayerlistRemove(conn.Info.UUID)
		//gui.RemovePlayer(cntrl.UUID)
	}
}

func (srv *Server) addPlayer(p *PlayerController) {
	srv.mu.Lock()
	srv.Players[p.UUID] = p
	srv.mu.Unlock()

	srv.PlayerlistUpdate()
	//gui.AddPlayer(p.session.Info().Name, p.UUID)

	srv.Logger.Info("[%s] Player %s (%s) has joined the server", p.session.RemoteAddr().String(), p.session.Info().Name, p.UUID)
}

func (srv *Server) GetCommand(name string) func(commands.Executor, []string) {
	var cmd func(commands.Executor, []string)
	for _, c := range srv.CommandGraph.Commands {
		if c.Name == name {
			return c.Execute
		}

		for _, a := range c.Aliases {
			if a == name {
				return c.Execute
			}
		}
	}

	return cmd
}
