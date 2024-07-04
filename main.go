package main

import (
	"aether/server"
	"net"
)

func main() {
	cfg := server.ServerConfig{
		IP:                   net.IPv4(127, 0, 0, 1),
		Port:                 25565,
		TPS:                  20,
		CompressionThreshold: -1,
	}
	srv, _ := cfg.New()
	srv.Start()
}
