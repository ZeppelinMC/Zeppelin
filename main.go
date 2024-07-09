package main

import (
	"aether/log"
	"aether/net/registry"
	"aether/server"
	"aether/server/world/region/blocks"
	"net"
	"runtime"
	"time"
)

var timeStart = time.Now()

func main() {
	log.Infoln("Aether 1.21 Minecraft server")
	log.Infof("Running on platform %s-%s\n", runtime.GOOS, runtime.GOARCH)
	log.Infoln("Loading embedded 1.21 server registries")
	if err := registry.LoadRegistry(); err != nil {
		log.Errorln("Error loading server registries:", err)
		return
	}
	if err := blocks.LoadBlockCache(); err != nil {
		log.Errorln("Error loading server registries:", err)
		return
	}

	var ip = net.IPv4(127, 0, 0, 1)
	var port = 25565

	log.Infof("Binding server to %s:%d TCP\n", ip, port)
	cfg := server.ServerConfig{
		IP:                   ip,
		Port:                 port,
		TPS:                  20,
		CompressionThreshold: -1,
	}
	srv, err := cfg.New()
	if err != nil {
		log.Errorln("Error binding server:", err)
		return
	}
	srv.Start(timeStart)
}
