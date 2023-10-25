package main

import (
	"bufio"
	"fmt"
	"github.com/dynamitemc/dynamite/server/commands"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"

	"github.com/pelletier/go-toml/v2"

	"github.com/dynamitemc/dynamite/core_commands"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/util"
	"github.com/dynamitemc/dynamite/web"
)

var log = logger.New()
var startTime = time.Now()

func startProfile() {
	file, _ := os.Create("cpu.out")
	pprof.StartCPUProfile(file)
}

func stopProfile() {
	pprof.StopCPUProfile()
	file, _ := os.Create("ram.out")
	runtime.GC()
	pprof.WriteHeapProfile(file)
	file.Close()
}

func start(cfg *server.Config) {
	srv, err := server.Listen(cfg, cfg.ServerIP+":"+strconv.Itoa(cfg.ServerPort), log, core_commands.Commands)
	log.Info("Opened TCP server on %s:%d", cfg.ServerIP, cfg.ServerPort)
	if err != nil {
		log.Error("Failed to open TCP server: %s", err)
		os.Exit(1)
	}
	log.Info("Done! (%v)", time.Since(startTime))
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		if util.HasArg("-prof") {
			stopProfile()
		}
		srv.Close()
	}()

	go scanConsole(srv)
	err = srv.Start()
	if err != nil {
		log.Error("Failed to start server: %s", err)
		os.Exit(1)
	}
}

var cfg server.Config

func main() {
	log.Info("Starting Dynamite 1.20.1 server")
	if util.HasArg("-prof") {
		log.Info("Starting CPU/RAM profiler")
		startProfile()
	}

	if err := server.LoadConfig("config.toml", &cfg); err != nil {
		log.Info("%v loading config.toml. Using default config", err)
		cfg = server.DefaultConfig

		f, _ := os.OpenFile("config.toml", os.O_RDWR|os.O_CREATE, 0666)
		toml.NewEncoder(f).Encode(cfg)
	}
	log.Debug("Loaded config")

	if !cfg.Online && !util.HasArg("-no_offline_warn") {
		log.Warn("Offline mode is insecure and you should not use it unless for a private server.\nRead https://github.com/DynamiteMC/Dynamite/wiki/Why-you-shouldn't-use-offline-mode")
	}

	if cfg.Web.Enable {
		if !util.HasArg("-nogui") {
			go web.LaunchWebPanel(fmt.Sprintf("%s:%d", cfg.Web.ServerIP, cfg.Web.ServerPort), cfg.Web.Password, log)
		} else {
			log.Warn("Remove the -nogui argument to load the web panel")
		}
	}
	start(&cfg)
}

func scanConsole(srv *server.Server) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			continue
		}

		content := strings.TrimSpace(txt)
		args := strings.Split(content, " ")

		command := srv.GetCommandGraph().FindCommand(args[0])
		if command == nil {
			srv.Logger.Print(fmt.Sprintf("&cUnknown or incomplete command, see below for error\n&n%s&r&c&o<--[HERE]", args[0]))
			continue
		}
		command.Execute(commands.CommandContext{
			Arguments:   args[1:],
			Executor:    &server.ConsoleExecutor{Server: srv},
			FullCommand: content,
		})
	}

	if err := scanner.Err(); err != nil {
		srv.Logger.Error("%v scanning console", err)
	}
}
