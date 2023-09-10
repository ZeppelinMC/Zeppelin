package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dynamitemc/dynamite/config"
	"github.com/dynamitemc/dynamite/gui"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/util"
)

var log logger.Logger
var startTime = time.Now().Unix()

func start(cfg server.ServerConfig) {
	srv, err := cfg.Listen(cfg.ServerIP+":"+strconv.Itoa(cfg.ServerPort), log)
	log.Info("Opened TCP server on %s:%d", cfg.ServerIP, cfg.ServerPort)
	if err != nil {
		log.Error("Failed to open TCP server: %s", err)
		os.Exit(1)
	}
	log.Info("Done! (%ds)", time.Now().Unix()-startTime)
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
	if cfg.Web.Enable {
		if !util.HasArg("-nogui") {
			go gui.LaunchWebPanel(fmt.Sprintf("%s:%d", cfg.Web.ServerIP, cfg.Web.ServerPort), cfg.Web.Password, &log)
		} else {
			log.Warn("Remove the -nogui argument to load the web panel")
		}
	}
	if cfg.GUI {
		if !util.HasArg("-nogui") {
			go start(cfg)
			log.Info("Loading GUI panel")
			gui.LaunchGUI(&log)
		} else {
			log.Warn("Remove the -nogui argument to load the GUI panel")
			start(cfg)
		}
	} else {
		start(cfg)
	}
}
