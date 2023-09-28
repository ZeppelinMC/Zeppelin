package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/dynamitemc/dynamite/config"
	"github.com/dynamitemc/dynamite/core_commands"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/util"
	"github.com/dynamitemc/dynamite/web"
)

var log logger.Logger
var startTime = time.Now().Unix()

func start(cfg server.ServerConfig) {
	srv, err := cfg.Listen(cfg.ServerIP+":"+strconv.Itoa(cfg.ServerPort), log, core_commands.Commands)
	log.Info("Opened TCP server on %s:%d", cfg.ServerIP, cfg.ServerPort)
	if err != nil {
		log.Error("Failed to open TCP server: %s", err)
		os.Exit(1)
	}
	srv.LoadPlugins()
	log.Info("Done! (%ds)", time.Now().Unix()-startTime)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			srv.Close()
		}
	}()
	go srv.ScanConsole()
	err = srv.Start()
	if err != nil {
		log.Error("Failed to start server: %s", err)
		os.Exit(1)
	}
}

func main() {
	log.Info("Starting Dynamite Server")
	var cfg server.ServerConfig
	config.LoadConfig("config.toml", &cfg)
	log.Debug("Loaded config")

	if !cfg.Online && !util.HasArg("-no_offline_warn") {
		log.Warn("Offline mode is insecure and you should not use it unless for a private server.")
	}

	if cfg.Web.Enable {
		if !util.HasArg("-nogui") {
			go web.LaunchWebPanel(fmt.Sprintf("%s:%d", cfg.Web.ServerIP, cfg.Web.ServerPort), cfg.Web.Password, &log)
		} else {
			log.Warn("Remove the -nogui argument to load the web panel")
		}
	}
	start(cfg)
}
