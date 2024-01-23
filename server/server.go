package server

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/aimjel/minecraft"
	"github.com/aimjel/nitrate/server/network"
	"github.com/aimjel/nitrate/server/player"
	"github.com/aimjel/nitrate/server/world"
)

type Server struct {
	cfg *Config

	ln *minecraft.Listener

	Logger *log.Logger

	world *world.World

	broadcaster *network.Broadcast
}

func (cfg *Config) NewServer(status *minecraft.Status) (*Server, error) {
	ln, err := (&minecraft.ListenConfig{
		Status:               status,
		OnlineMode:           cfg.OnlineMode,
		CompressionThreshold: cfg.CompressionThreshold,
		Messages:             nil, //todo
	}).Listen(cfg.Address)

	if err != nil {
		return nil, err
	}

	var w *world.World
	//todo change to use a variable
	if w, err = world.OpenWorld("world2"); err != nil {
		return nil, err
	}

	srv := &Server{
		cfg:         cfg,
		ln:          ln,
		world:       w,
		broadcaster: network.NewBroadcast(),
	}

	srv.handleSIGINT()

	return srv, nil
}

func (srv *Server) Accept() (*player.Player, error) {
	conn, err := srv.ln.Accept()
	if err != nil {
		return nil, err
	}
	plyer := player.New(srv.world)
	go srv.handleNewConn(plyer, conn)
	return plyer, err
}

func (srv *Server) handleNewConn(p *player.Player, c *minecraft.Conn) {
	sesh := network.NewSession(c, p, srv.broadcaster)

	p.Session = sesh

	sesh.ViewDistance = srv.cfg.ViewDistance
	sesh.LoginPlay()

	srv.broadcaster.AddSession(sesh)

	if err := sesh.HandlePackets(); err != nil {
		srv.broadcaster.RemoveSessions(sesh)
		//error occurred while handling packets
		srv.Logger.Println(err)
	}
}

func (srv *Server) handleSIGINT() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			fmt.Println("you quit! lol")
			os.Exit(0)
		}
	}()
}
