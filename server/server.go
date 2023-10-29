package server

import (
	"errors"
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/aimjel/minecraft/chat"
	"github.com/aimjel/minecraft/protocol/types"
	"github.com/pelletier/go-toml/v2"

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
	players map[string]*Session

	WhitelistedPlayers,
	Operators,
	BannedPlayers,
	BannedIPs []user

	listener *minecraft.Listener

	entities map[int32]*Entity

	World *world.World

	mu sync.RWMutex
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
		conn.Close(nil)
		return
	}

	x := conn.UUID()
	uuid, _ := uuid.FromBytes(x[:])

	data := srv.World.GetPlayerData(uuid.String())

	plyr := player.New(data)
	cntrl := &Session{
		Player:   plyr,
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

	//cntrl.SendPacket(&packet.SetTablistHeaderFooter{
	//	Header: strings.Join(srv.Config.Tablist.Header, "\n"),
	//	Footer: strings.Join(srv.Config.Tablist.Footer, "\n"),
	//})

	srv.addPlayer(cntrl)
	cntrl.Login(plyr.Dimension())

	cntrl.intitializeData()

	if err := cntrl.HandlePackets(); err != nil {
		prefix, suffix := cntrl.GetPrefixSuffix()
		srv.Logger.Info("[%s] Player %s (%s) has left the server", conn.RemoteAddr().String(), conn.Name(), cntrl.UUID)

		srv.GlobalMessage(srv.Translate("multiplayer.player.left", chat.NewMessage(prefix+conn.Name()+suffix)))
		srv.PlayerlistRemove(conn.UUID())
		cntrl.Despawn()
		plyr.Save()

		//todo consider moving logic of removing player to a separate function
		srv.mu.Lock()
		delete(srv.players, cntrl.UUID)
		srv.mu.Unlock()
		//gui.RemovePlayer(cntrl.UUID)
	}
}

func (srv *Server) addPlayer(p *Session) {
	srv.mu.Lock()
	srv.players[p.UUID] = p
	srv.mu.Unlock()
	newPlayer := types.PlayerInfo{
		UUID:       p.conn.UUID(),
		Name:       p.conn.Name(),
		Properties: p.conn.Properties(),
		Listed:     true,
	}

	srv.mu.RLock()

	players := make([]types.PlayerInfo, 0, len(srv.players))
	for _, pl := range srv.players {
		players = append(players, types.PlayerInfo{
			UUID:          pl.conn.UUID(),
			Name:          pl.conn.Name(),
			Properties:    pl.conn.Properties(),
			Listed:        true,
			PublicKey:     pl.publicKey,
			KeySignature:  pl.keySignature,
			ChatSessionID: pl.sessionID,
			ExpiresAt:     int64(pl.expires),
		})
		pl.SendPacket(&packet.PlayerInfoUpdate{
			Actions: 0x01 | 0x08,
			Players: []types.PlayerInfo{newPlayer},
		})
	}
	srv.mu.RUnlock()

	//updates the new session's player list
	p.SendPacket(&packet.PlayerInfoUpdate{
		Actions: 0x01 | 0x02 | 0x08,
		Players: players,
	})

	//gui.AddPlayer(pl.session.Info().Name, pl.UUID)
	prefix, suffix := p.GetPrefixSuffix()

	srv.Logger.Info("[%s] Player %s (%s) has joined the server", p.conn.RemoteAddr(), p.conn.Name(), p.UUID)

	srv.GlobalMessage(srv.Translate("multiplayer.player.joined", chat.NewMessage(prefix+p.Name()+suffix)))
}

func (srv *Server) GetCommandGraph() *commands.Graph {
	return srv.commandGraph
}

func (srv *Server) Translate(msg string, with ...chat.Message) chat.Message {
	return chat.Translate(msg, with...)
}

func (srv *Server) ParsePlaceholders(msg string, data map[string]string) string {
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

	for _, p := range srv.players {
		if srv.Config.Whitelist.Enforce && srv.Config.Whitelist.Enable && !srv.IsWhitelisted(p.conn.UUID()) {
			p.Disconnect(chat.Translate("multiplayer.disconnect.not_whitelisted"))
			continue
		}

		p.Player.SetOperator(srv.IsOperator(p.conn.UUID()))

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
		return srv.entities[id]
	}
}

func (srv *Server) FindEntityByUUID(id [16]byte) interface{} {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	for _, p := range srv.players {
		if p.conn.UUID() == id {
			return p
		}
	}
	for _, e := range srv.entities {
		if e.UUID == id {
			return e
		}
	}
	return nil
}

func (srv *Server) Player(uuid string) *Session {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	return srv.players[uuid]
}

func (srv *Server) PlayerCount() int {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	return len(srv.players)
}

func (srv *Server) Players() map[string]*Session {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	return srv.players
}

func (srv *Server) FindPlayer(username string) *Session {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	for _, p := range srv.players {
		if strings.EqualFold(p.Name(), username) {
			return p
		}
	}
	return nil
}

func (srv *Server) FindPlayerByID(id int32) *Session {
	srv.mu.RLock()
	defer srv.mu.RUnlock()
	for _, p := range srv.players {
		if p.entityID == id {
			return p
		}
	}
	return nil
}

func (srv *Server) Close() {
	srv.mu.Lock()
	defer srv.mu.Unlock()
	srv.Logger.Info("Closing server...")

	var files = []string{"whitelist.json", "banned_players.json", "ops.json", "banned_ips.json"}
	var lists = [][]user{srv.WhitelistedPlayers, srv.BannedPlayers, srv.Operators, srv.BannedIPs}
	for i, file := range files {
		WritePlayerList(file, lists[i])
	}

	saveCache()

	for _, p := range srv.players {
		p.Disconnect(srv.Translate("multiplayer.disconnect.server_shutdown"))
		p.Player.Save()
	}
	srv.Logger.Info("Saving world...")
	srv.World.Save()

	f, _ := os.OpenFile("config.toml", os.O_RDWR|os.O_CREATE, 0666)
	_ = toml.NewEncoder(f).Encode(srv.Config)
	os.Exit(0)
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
