package server

import (
	"strings"
	"sync"

	"github.com/aimjel/minecraft"
	"github.com/aimjel/minecraft/chat"
	"github.com/aimjel/minecraft/packet"
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

	WhitelistedPlayers []util.Player
	Operators          []util.Player
	BannedPlayers      []util.Player
	BannedIPs          []string

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
	var reason string
	if r := srv.ValidatePlayer(session.Conn.Info.Name, player.UUID, strings.Split(session.Conn.RemoteAddr().String(), ":")[0]); r != CONNECTION_VALID {
		switch r {
		case CONNECTION_SERVER_FULL:
			{
				reason = srv.Config.Messages.ServerFull
			}
		case CONNECTION_PLAYER_BANNED:
			{
				reason = srv.Config.Messages.Banned
			}
		case CONNECTION_PLAYER_ALREADY_PLAYING:
			{
				reason = srv.Config.Messages.AlreadyPlaying
			}
		case CONNECTION_PLAYER_NOT_IN_WHITELIST:
			{
				reason = srv.Config.Messages.NotInWhitelist
			}
		}
		msg := chat.NewMessage(reason)
		conn.SendPacket(&packet.DisconnectLogin{Reason: msg.String()})
		return
	}
	srv.addPlayer(player)
	player.SetCommands(srv.CommandGraph.Data())
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

		srv.Logger.Info("[%s] Player %s (%s) has left the server", conn.RemoteAddr().String(), conn.Info.Name, player.UUID)
		srv.PlayerlistRemove(u)
		gui.RemovePlayer(player.UUID)
	}
}

func (srv *Server) addPlayer(p *p.Player) {
	srv.Lock()
	srv.Players[p.UUID] = p
	srv.Unlock()
	srv.PlayerlistUpdate()
	gui.AddPlayer(p.Session.Conn.Info.Name, p.UUID)

	srv.Logger.Info("[%s] Player %s (%s) has joined the server", p.Session.Conn.RemoteAddr().String(), p.Session.Conn.Info.Name, p.UUID)
}

func (srv *Server) GetCommand(name string) func(*p.Player, []string) {
	var cmd func(*p.Player, []string)
	for _, c := range srv.CommandGraph.Commands {
		if c.Name == name {
			return c.Execute
		} else {
			for _, a := range c.Aliases {
				if a == name {
					return c.Execute
				}
			}
		}
	}
	return cmd
}
