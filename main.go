package main

import (
	"aether/net/registry"
	"aether/server"
	"fmt"
	"net"
)

func main() {
	fmt.Print("Loading 1.21 server registries: ")
	fmt.Println(registry.LoadRegistry())

	cfg := server.ServerConfig{
		IP:                   net.IPv4(127, 0, 0, 1),
		Port:                 25565,
		TPS:                  20,
		CompressionThreshold: -1,
	}
	srv, _ := cfg.New()
	srv.Start()
}
