package server

import (
	"aether/net"
	"aether/net/packet/configuration"
	"fmt"
)

type Server struct {
	cfg      ServerConfig
	listener *net.Listener
	ticker   Ticker
}

func (srv Server) Start() {
	fmt.Println("started ticker!")
	srv.ticker.Start()
	fmt.Println("started server")
	for {
		conn, err := srv.listener.Accept()
		if err != nil {
			fmt.Println("server error", err)
			return
		}
		if conn == nil {
			continue
		}
		fmt.Println("new connection from player", conn.Username())
		conn.SetState(net.PlayState)
		conn.WritePacket(configuration.FinishConfiguration{})
	}
}
