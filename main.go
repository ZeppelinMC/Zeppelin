package main

import (
	"strconv"
	"time"

	"github.com/dynamitemc/dynamite/config"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/server/commands"
)

var log logger.Logger
var startTime = time.Now().Unix()

func main() {
	log.Info("Starting DynamiteMC")
	var cfg server.ServerConfig
	config.LoadConfig("config.toml", &cfg)
	log.Debug("Loaded config")

	server, err := cfg.Listen(cfg.ServerIP+":"+strconv.Itoa(cfg.ServerPort), log)
	log.Info("Opened TCP server on %s:%d", cfg.ServerIP, cfg.ServerPort)
	if err != nil {
		panic(err)
	}
	server.CommandGraph.AddCommands(commands.NewCommand("hi"))
	log.Info("Done! (%ds)", time.Now().Unix()-startTime)
	server.Start()
}
