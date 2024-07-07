package main

import (
	"aether/nbt"
	"aether/net/registry"
	"aether/server"
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Print("Loading 1.21 server registries: ")
	fmt.Println(registry.LoadRegistry())

	f, _ := os.Open(`C:\Users\oqammx86\Downloads\Untitled1.nbt`)
	d := nbt.NewDecoder(f)
	d.ReadRootName(false)
	var data map[string]any
	fmt.Println(d.Decode(&data))
	fmt.Println(data)

	cfg := server.ServerConfig{
		IP:                   net.IPv4(127, 0, 0, 1),
		Port:                 25565,
		TPS:                  20,
		CompressionThreshold: -1,
	}
	srv, _ := cfg.New()
	srv.Start()
}
