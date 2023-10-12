package server

import (
	"errors"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/aimjel/minecraft"

	//"github.com/dynamitemc/dynamite/web"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/player"
	"github.com/dynamitemc/dynamite/server/world"
)

type Server struct {
	Config       *Config
	Logger       *logger.Logger
	CommandGraph *commands.Graph

	Plugins map[string]*Plugin

	// Players mapped by UUID
	Players map[string]*PlayerController

	WhitelistedPlayers,
	Operators,
	BannedPlayers,
	BannedIPs []user

	listener *minecraft.Listener

	teleportCounter int32

	entityCounter int32

	World *world.World

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

	uuid, _ := uuid.FromBytes(conn.Info.UUID[:])

	data := srv.World.GetPlayerData(uuid.String())

	plyr := player.New(srv.entityCounter, int32(srv.Config.ViewDistance), int32(srv.Config.SimulationDistance), data)
	sesh := New(conn, plyr)
	cntrl := &PlayerController{player: plyr, session: sesh, Server: srv}
	cntrl.UUID = uuid.String()

	for _, op := range srv.Operators {
		if op.UUID == cntrl.UUID {
			plyr.SetOperator(true)
		}
	}

	cntrl.SendCommands(srv.CommandGraph)

	srv.addPlayer(cntrl)
	if err := cntrl.Login(srv.World.Overworld()); err != nil {
		//TODO log error
		conn.Close(err)
		srv.Logger.Error("Failed to join player to dimension %s", err)
	}

	go func() {
		ticker := time.NewTicker(10 * time.Second)
		for range ticker.C {
			cntrl.Keepalive()
		}
	}()

	cntrl.InitializeInventory()

	if err := sesh.HandlePackets(cntrl); err != nil {
		srv.Logger.Info("[%s] Player %s (%s) has left the server", conn.RemoteAddr().String(), conn.Info.Name, cntrl.UUID)
		srv.GlobalMessage(srv.Translate(srv.Config.Messages.PlayerLeave, map[string]string{"player": conn.Info.Name}), nil)
		srv.PlayerlistRemove(conn.Info.UUID)
		cntrl.Despawn()

		//todo consider moving logic of removing player to a separate function
		srv.mu.Lock()
		delete(srv.Players, cntrl.UUID)
		srv.mu.Unlock()
		//gui.RemovePlayer(cntrl.UUID)
	}
}

func (srv *Server) addPlayer(p *PlayerController) {
	srv.mu.RLock()
	srv.Players[p.UUID] = p
	srv.mu.RUnlock()

	srv.PlayerlistUpdate()
	//gui.AddPlayer(p.session.Info().Name, p.UUID)

	srv.Logger.Info("[%s] Player %s (%s) has joined the server", p.session.RemoteAddr().String(), p.session.Info().Name, p.UUID)
	srv.GlobalMessage(srv.Translate(srv.Config.Messages.PlayerJoin, map[string]string{"player": p.Name()}), nil)
}

func (srv *Server) GetCommandGraph() *commands.Graph {
	return srv.CommandGraph
}

func (srv *Server) Translate(msg string, data map[string]string) string {
	for k, v := range data {
		msg = strings.ReplaceAll(msg, "%"+k+"%", v)
	}
	return msg
}

func (srv *Server) Reload() error {
	srv.loadFiles()

	LoadConfig("config.toml", srv.Config)

	clearCache()

	srv.mu.RLock()
	defer srv.mu.RUnlock()

	for _, p := range srv.Players {
		if srv.Config.Whitelist.Enforce && srv.Config.Whitelist.Enable && !srv.IsWhitelisted(p.session.Info().UUID) {
			p.Disconnect(srv.Config.Messages.NotInWhitelist)
			continue
		}

		p.player.SetOperator(srv.IsOperator(p.session.Info().UUID))

		p.SendCommands(srv.CommandGraph)
	}
	return nil
}

func (srv *Server) FindPlayer(username string) *PlayerController {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	for _, p := range srv.Players {
		if strings.EqualFold(p.Name(), username) {
			return p
		}
	}
	return nil
}

func (srv *Server) FindPlayerByID(id int32) *PlayerController {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	for _, p := range srv.Players {
		if p.player.EntityId() == id {
			return p
		}
	}
	return nil
}

func (srv *Server) Close() {
	srv.Logger.Info("Closing server...")

	var files = []string{"whitelist.json", "banned_players.json", "ops.json", "banned_ips.json"}
	var lists = [][]user{srv.WhitelistedPlayers, srv.BannedPlayers, srv.Operators, srv.BannedIPs}
	for i, file := range files {
		WritePlayerList(file, lists[i])
	}

	saveCache()

	for _, p := range srv.Players {
		p.player.Save()
		p.Disconnect(srv.Config.Messages.ServerClosed)
	}
}

func (srv *Server) loadFiles() {
	var files = []string{"whitelist.json", "banned_players.json", "ops.json", "banned_ips.json"}
	var addresses = []*[]user{&srv.WhitelistedPlayers, &srv.BannedPlayers, &srv.Operators, &srv.BannedIPs}
	for i, file := range files {
		u, err := loadUsers(file)
		if err != nil && !errors.Is(err, os.ErrNotExist) {
			srv.Logger.Warn("%v loading %v", err, file)
		}

		*addresses[i] = u
	}
}
