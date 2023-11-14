package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
	"time"

	"github.com/pelletier/go-toml/v2"
	"golang.org/x/term"

	"github.com/dynamitemc/dynamite/core_commands"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/server/config"
	"github.com/dynamitemc/dynamite/server/world/tick"
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

func start(cfg *config.Config) {
	srv, err := server.New(cfg, cfg.ServerIP+":"+strconv.Itoa(cfg.ServerPort), log, core_commands.Commands)
	log.Info("Opened TCP server on %s:%d", cfg.ServerIP, cfg.ServerPort)
	if err != nil {
		log.Error("Failed to open TCP server: %s", err)
		os.Exit(1)
	}
	ticker := tick.New(srv, srv.Config.TPS)
	ticker.Start()
	log.Info("Started tick loop")
	log.Info("Done! (%v)", time.Since(startTime))

	go scanConsole(srv)
	defer stopProfile()
	err = srv.Start()
	if err != nil {
		log.Error("Failed to start server: %s", err)
		os.Exit(1)
	}
}

var cfg config.Config

func main() {
	server.OldState, _ = term.MakeRaw(int(os.Stdin.Fd()))

	log.Info("Starting Dynamite 1.20.1 server")
	if util.HasArg("-prof") {
		log.Info("Starting CPU/RAM profiler")
		startProfile()
	}

	if err := config.LoadConfig("config.toml", &cfg); err != nil {
		log.Info("%v loading config.toml. Using default config", err)
		cfg = config.DefaultConfig

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

// The extremely fancy custom terminal thing
func scanConsole(srv *server.Server) {
	var command string
	var b [1]byte
	for {
		os.Stdin.Read(b[:])

		switch b[0] {
		case 127, 8: // backspace - delete character
			if len(command) > 0 {
				fmt.Print("\b \b")
				command = command[:len(command)-1]
				args := strings.Split(command, " ")

				cmd := srv.GetCommandGraph().FindCommand(args[0])
				if cmd == nil {
					fmt.Printf("\r> %s", logger.HR(command))
				} else {
					if len(args) > 1 {
						fmt.Printf("\r> %s %s", args[0], logger.C(strings.Join(args[1:], " ")))
					} else {
						fmt.Printf("\r> %s", args[0])
					}
				}
			}
		case 3: // ctrl c - stop the server
			fmt.Print("\r")
			if len(command) > len("stop") {
				fmt.Print("\x1b[K")
			}
			fmt.Print("> ")
			srv.ConsoleCommand("stop")
		case 13: // enter - run the command and clear it
			command = strings.TrimSpace(command)
			if command == "" {
				continue
			}
			fmt.Printf("\r> %s\r> ", strings.Repeat(" ", len(command)))
			srv.ConsoleCommand(command)
			command = ""
		case 65, 27, 66, 91, 67, 68:
		default: // regular character - add to current command input
			command += string(b[0])
			args := strings.Split(command, " ")

			cmd := srv.GetCommandGraph().FindCommand(args[0])
			if cmd == nil {
				fmt.Printf("\r> %s", logger.HR(command))
			} else {
				if len(args) > 1 {
					fmt.Printf("\r> %s %s", args[0], logger.C(strings.Join(args[1:], " ")))
				} else {
					fmt.Printf("\r> %s", args[0])
				}
			}
		}
	}
}
