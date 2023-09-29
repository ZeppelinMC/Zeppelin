package server

import (
	"crypto/md5"
	"errors"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/aimjel/minecraft"

	//"github.com/dynamitemc/dynamite/web"
	"github.com/dynamitemc/dynamite/config"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/player"
	"github.com/dynamitemc/dynamite/server/plugins"
	"github.com/dynamitemc/dynamite/server/world"
)

type Server struct {
	Config       *config.ServerConfig
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

func NameToUUID(name string) uuid.UUID {
	version := 3
	h := md5.New()
	h.Write([]byte("OfflinePlayer:"))
	h.Write([]byte(name))
	var id uuid.UUID
	h.Sum(id[:0])
	id[6] = (id[6] & 0x0f) | uint8((version&0xf)<<4)
	id[8] = (id[8] & 0x3f) | 0x80 // RFC 4122 variant
	return id
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

	//cntrl.SendCommands(srv.CommandGraph)

	srv.addPlayer(cntrl)
	if err := cntrl.JoinDimension(srv.world.DefaultDimension()); err != nil {
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
	if err := sesh.HandlePackets(cntrl); err != nil {
		srv.Logger.Info("[%s] Player %s (%s) has left the server", conn.RemoteAddr().String(), conn.Info.Name, cntrl.UUID)
		srv.GlobalMessage(srv.Translate(srv.Config.Messages.PlayerLeave, map[string]string{"player": conn.Info.Name}))
		srv.PlayerlistRemove(conn.Info.UUID)
		delete(srv.Players, cntrl.UUID)
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
	srv.GlobalMessage(srv.Translate(srv.Config.Messages.PlayerJoin, map[string]string{"player": p.Name()}))
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

	// load config
	config.LoadConfig("config.toml", srv.Config)

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

func (srv *Server) Close() {
	srv.Logger.Info("Closing server...")
	os.Exit(0)
}
