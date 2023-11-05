package server

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/dynamitemc/dynamite/server/permission"

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
	"github.com/dynamitemc/dynamite/web"
)

var idCounter atomic.Int32

type Server struct {
	Config       *Config
	Logger       *logger.Logger
	commandGraph *commands.Graph

	lang map[string]string

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
	cntrl.uuid = uuid.String()
	cntrl.clientInfo.ViewDistance = int8(srv.Config.ViewDistance)

	for _, op := range srv.Operators {
		if op.UUID == cntrl.UUID() {
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

	if err := cntrl.handlePackets(); err != nil {
		prefix, suffix := cntrl.GetPrefixSuffix()
		srv.Logger.Info("[%s] Player %s (%s) has left the server", conn.RemoteAddr().String(), conn.Name(), cntrl.UUID())

		srv.GlobalMessage(srv.Translate("player.leave", map[string]string{"player": cntrl.Name(), "player_prefix": prefix, "player_suffix": suffix}))
		srv.PlayerlistRemove(conn.UUID())
		cntrl.Despawn()
		plyr.Save()

		//todo consider moving logic of removing player to a separate function
		srv.mu.Lock()
		delete(srv.players, cntrl.UUID())
		srv.mu.Unlock()
		//gui.RemovePlayer(cntrl.UUID)
	}
}

func (srv *Server) addPlayer(p *Session) {
	p.SendPacket(&packet.ServerData{
		MOTD:               chat.NewMessage(srv.Config.MOTD),
		EnforcesSecureChat: true,
	})
	newPlayer := types.PlayerInfo{
		UUID:       p.conn.UUID(),
		Name:       p.conn.Name(),
		Properties: p.conn.Properties(),
		Listed:     true,
	}

	srv.mu.Lock()
	srv.players[p.UUID()] = p
	players := make([]types.PlayerInfo, 0, len(srv.players))
	for _, pl := range srv.players {
		pl.mu.RLock()
		players = append(players, types.PlayerInfo{
			UUID:          pl.conn.UUID(),
			Name:          pl.conn.Name(),
			Properties:    pl.conn.Properties(),
			Listed:        true,
			PublicKey:     bytes.Clone(pl.publicKey),
			KeySignature:  bytes.Clone(pl.keySignature),
			ChatSessionID: pl.sessionID,
			ExpiresAt:     pl.expires,
		})
		if pl.UUID() != p.UUID() {
			pl.SendPacket(&packet.PlayerInfoUpdate{
				Actions: 0x01 | 0x08,
				Players: []types.PlayerInfo{newPlayer},
			})
		}
		pl.mu.RUnlock()
	}
	srv.mu.Unlock()

	//updates the new session's player list
	var actions byte = 0x01 | 0x08
	if p.Server.Config.Chat.Secure {
		actions |= 0x02
	}
	p.SendPacket(&packet.PlayerInfoUpdate{
		Actions: actions,
		Players: players,
	})

	if srv.Config.Web.Enable {
		web.AddPlayer(p.conn.Name(), p.UUID())
	}
	prefix, suffix := p.GetPrefixSuffix()

	srv.Logger.Info("[%s] Player %s (%s) has joined the server", p.conn.RemoteAddr(), p.conn.Name(), p.UUID())

	srv.GlobalMessage(srv.Translate("player.join", map[string]string{"player": p.Name(), "player_prefix": prefix, "player_suffix": suffix}))
}

func (srv *Server) GetCommandGraph() *commands.Graph {
	return srv.commandGraph
}

func (srv *Server) Translate(msg string, data map[string]string) chat.Message {
	txt, ok := srv.lang[msg]
	if !ok {
		return chat.NewMessage(msg)
	}
	return srv.ParsePlaceholders(txt, data)
}

func (srv *Server) ParsePlaceholders(txt string, data map[string]string) chat.Message {
	for k, v := range data {
		txt = strings.ReplaceAll(txt, "%"+k+"%", v)
	}
	return chat.NewMessage(txt)
}

func (srv *Server) Reload() error {
	srv.loadFiles()

	LoadConfig("config.toml", srv.Config)

	permission.Clear()

	srv.mu.RLock()
	defer srv.mu.RUnlock()

	for _, p := range srv.players {
		if srv.Config.Whitelist.Enforce && srv.Config.Whitelist.Enable && !srv.IsWhitelisted(p.conn.UUID()) {
			p.Disconnect(srv.Translate("disconnect.not_whitelisted", nil))
			continue
		}

		p.Player.SetOperator(srv.IsOperator(p.conn.UUID()))

		p.SendCommands(srv.commandGraph)
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

	permission.Save()

	for _, p := range srv.players {
		p.Disconnect(srv.Translate("disconnect.server_shutdown", nil))
		p.Player.Save()
	}
	srv.Logger.Done()
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

func (srv *Server) ConsoleCommand(txt string) {
	fmt.Println(txt)
	fmt.Print("\r")
	content := strings.TrimSpace(txt)
	args := strings.Split(content, " ")

	command := srv.GetCommandGraph().FindCommand(args[0])
	if command == nil {
		srv.Logger.Print(chat.NewMessage(fmt.Sprintf("&cUnknown or incomplete command, see below for error\n&n%s&r&c&o<--[HERE]", args[0])))
		return
	}
	command.Execute(commands.CommandContext{
		Command:     command,
		Arguments:   args[1:],
		Executor:    &ConsoleExecutor{Server: srv},
		FullCommand: content,
	})
}

type ConsoleExecutor struct {
	Server *Server
}
