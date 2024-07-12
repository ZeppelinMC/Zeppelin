package main

import (
	"runtime"
	"time"

	"github.com/dynamitemc/aether/log"
	"github.com/dynamitemc/aether/net/registry"
	"github.com/dynamitemc/aether/server/world/region/blocks"
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
	log.Infoln("Loading config")
	cfg := loadConfig()

	log.Infof("Binding server to %s:%d TCP\n", cfg.ServerIP, cfg.ServerPort)
	srv, err := cfg.New()
	if err != nil {
		log.Errorln("Error binding server:", err)
		return
	}
	srv.Start(timeStart)
}
