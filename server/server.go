package server

import (
	"aether/net"
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
		fmt.Println(srv.listener.Accept())
	}
}
