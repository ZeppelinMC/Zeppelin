package server

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/google/uuid"

	"github.com/aimjel/minecraft"
	"github.com/aimjel/minecraft/packet"

	//"github.com/dynamitemc/dynamite/web"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/player"
	"github.com/dynamitemc/dynamite/server/world"
)

var idCounter atomic.Int32

type Server struct {
	Config       *Config
	Logger       *logger.Logger
	commandGraph *commands.Graph

	// Players mapped by UUID
	Players map[string]*Session

	WhitelistedPlayers,
	Operators,
	BannedPlayers,
	BannedIPs []user

	listener *minecraft.Listener

	Entities map[int32]*Entity

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

func (srv *Server) GetDimension(typ string) *world.Dimension {
	switch typ {
	case "minecraft:the_nether":
		return srv.World.Nether()
	case "minecraft:the_end":
		return srv.World.TheEnd()
	default:
		return srv.World.Overworld()
	}
}

func (srv *Server) handleNewConn(conn *minecraft.Conn) {
	if srv.ValidateConn(conn) {
		return
	}

	x := conn.UUID()
	uuid, _ := uuid.FromBytes(x[:])

	data := srv.World.GetPlayerData(uuid.String())

	plyr := player.New(data)
	cntrl := &Session{
		player:   plyr,
		conn:     conn,
		Server:   srv,
		entityID: idCounter.Add(1),
	}
	cntrl.UUID = uuid.String()
	cntrl.clientInfo.ViewDistance = int8(srv.Config.ViewDistance)

	for _, op := range srv.Operators {
		if op.UUID == cntrl.UUID {
			plyr.SetOperator(true)
		}
	}

	cntrl.SendCommands(srv.commandGraph)

	cntrl.SendPacket(&packet.SetTablistHeaderFooter{
		Header: strings.Join(srv.Config.Tablist.Header, "\n"),
		Footer: strings.Join(srv.Config.Tablist.Footer, "\n"),
	})

	fmt.Println("e")

	srv.addPlayer(cntrl)
	if err := cntrl.Login(plyr.Dimension()); err != nil {
		//TODO log error
		conn.Close(err)
		srv.Logger.Error("Failed to join player to dimension %s", err)
	}

	cntrl.InitializeInventory()

	if err := cntrl.HandlePackets(); err != nil {
		srv.Logger.Info("[%s] Player %s (%s) has left the server", conn.RemoteAddr().String(), conn.Name(), cntrl.UUID)
		srv.GlobalMessage(srv.Translate(srv.Config.Messages.PlayerLeave, map[string]string{"player": conn.Name()}), nil)
		srv.PlayerlistRemove(conn.UUID())
		cntrl.Despawn()
		plyr.Save()

		//todo consider moving logic of removing player to a separate function
		srv.mu.Lock()
		delete(srv.Players, cntrl.UUID)
		srv.mu.Unlock()
		//gui.RemovePlayer(cntrl.UUID)
	}
}

func (srv *Server) addPlayer(p *Session) {
	srv.mu.Lock()
	srv.Players[p.UUID] = p
	srv.mu.Unlock()

	srv.PlayerlistUpdate()

	//gui.AddPlayer(p.session.Info().Name, p.UUID)

	srv.Logger.Info("[%s] Player %s (%s) has joined the server", p.conn.RemoteAddr(), p.conn.Name(), p.UUID)
	srv.GlobalMessage(srv.Translate(srv.Config.Messages.PlayerJoin, map[string]string{"player": p.Name()}), nil)
}

func (srv *Server) GetCommandGraph() *commands.Graph {
	return srv.commandGraph
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
		if srv.Config.Whitelist.Enforce && srv.Config.Whitelist.Enable && !srv.IsWhitelisted(p.conn.UUID()) {
			p.Disconnect(srv.Config.Messages.NotInWhitelist)
			continue
		}

		p.player.SetOperator(srv.IsOperator(p.conn.UUID()))

		p.SendCommands(srv.commandGraph)
	}
	return nil
}

func (srv *Server) FindEntity(id int32) interface{} {
	if p := srv.FindPlayerByID(id); p != nil {
		return p
	} else {
		srv.mu.RLock()
		defer srv.mu.RUnlock()
		return srv.Entities[id]
	}
}

func (srv *Server) FindEntityByUUID(id [16]byte) interface{} {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	for _, p := range srv.Players {
		if p.conn.UUID() == id {
			return p
		}
	}
	for _, e := range srv.Entities {
		if e.UUID == id {
			return e
		}
	}
	return nil
}

func (srv *Server) FindPlayer(username string) *Session {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	for _, p := range srv.Players {
		if strings.EqualFold(p.Name(), username) {
			return p
		}
	}
	return nil
}

func (srv *Server) FindPlayerByID(id int32) *Session {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	for _, p := range srv.Players {
		if p.entityID == id {
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
		p.Disconnect(srv.Config.Messages.ServerClosed)
		p.player.Save()
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
