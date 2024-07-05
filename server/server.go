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
		conn, err := srv.listener.Accept()
		if err != nil {
			fmt.Println("server error", err)
			return
		}
		if conn == nil {
			continue
		}
	}
}
