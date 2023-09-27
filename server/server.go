package server

import (
	"errors"
	"os"
	"sync"

	"github.com/google/uuid"

	"github.com/aimjel/minecraft"

	//"github.com/dynamitemc/dynamite/web"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/player"
	"github.com/dynamitemc/dynamite/server/plugins"
	"github.com/dynamitemc/dynamite/server/world"
)

type Server struct {
	Config       *ServerConfig
	Logger       logger.Logger
	CommandGraph *commands.Graph

	Plugins map[string]*plugins.Plugin

	// Players mapped by UUID
	Players map[string]*PlayerController

	WhitelistedPlayers,
	Operators,
	BannedPlayers,
	BannedIPs []user

	listener *minecraft.Listener

	teleportCounter int32

	entityCounter int32

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
	srv.entityCounter++

	plyr := player.New(srv.entityCounter)
	sesh := New(conn, plyr)
	cntrl := &PlayerController{player: plyr, session: sesh, Server: srv}
	uuid, _ := uuid.FromBytes(conn.Info.UUID[:])
	cntrl.UUID = uuid.String()

	for _, op := range srv.Operators {
		if op.UUID == cntrl.UUID {
			plyr.Operator = true
		}
	}

	cntrl.SendCommands(srv.CommandGraph)

	if err := cntrl.JoinDimension(srv.world.DefaultDimension()); err != nil {
		//TODO log error
		conn.Close(err)
		srv.Logger.Error("Failed to join player to dimension %s", err)
	}

	srv.addPlayer(cntrl)
	if err := sesh.HandlePackets(cntrl); err != nil {
		srv.Logger.Info("[%s] Player %s (%s) has left the server", conn.RemoteAddr().String(), conn.Info.Name, cntrl.UUID)
		srv.PlayerlistRemove(conn.Info.UUID)
		delete(srv.Players, cntrl.UUID)
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

func (srv *Server) GetCommandGraph() *commands.Graph {
	return srv.CommandGraph
}

func (srv *Server) Reload() error {
	// load player data
	var files = []string{"whitelist.json", "banned_players.json", "ops.json", "banned_ips.json"}
	var addresses = []*[]user{&srv.WhitelistedPlayers, &srv.BannedPlayers, &srv.Operators, &srv.BannedIPs}
	for i, file := range files {
		u, err := loadUsers(file)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}

		*addresses[i] = u
	}
	for _, p := range srv.Players {
		p.SendCommands(srv.CommandGraph)
	}
	return nil
}

func (srv *Server) FindPlayer(username string) *PlayerController {
	for _, p := range srv.Players {
		if p.Name() == username {
			return p
		}
	}
	return nil
}
