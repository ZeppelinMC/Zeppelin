package main

import (
	"os"
	"runtime"
	"runtime/pprof"
	"time"

	"github.com/zeppelinmc/zeppelin/core_commands"
	"github.com/zeppelinmc/zeppelin/log"
	"github.com/zeppelinmc/zeppelin/server"
	"github.com/zeppelinmc/zeppelin/server/command"
	"github.com/zeppelinmc/zeppelin/server/world"
	"github.com/zeppelinmc/zeppelin/util"
	"golang.org/x/term"
)

var timeStart = time.Now()

func main() {
	log.Infolnf("Zeppelin 1.21 Minecraft server with %s on platform %s-%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)

	if util.HasArgument("--cpuprof") {
		f, _ := os.Create("zeppelin-cpu-profile")
		pprof.StartCPUProfile(f)
		log.Infoln("Started CPU profiler (writing to zeppelin-cpu-profile)")
	}

	cfg := loadConfig()

	w := world.NewWorld(cfg.LevelName)

	log.Infof("Binding server to %s:%d\n", cfg.ServerIp, cfg.ServerPort)

	rawTerminal := !util.HasArgument("--no-raw-terminal")

	srv, err := server.New(cfg, w)
	if err != nil {
		log.Errorln("Error binding server:", err)
		return
	}
	srv.CommandManager = command.NewManager(srv, core_commands.Commands...)
	var oldState *term.State
	if rawTerminal {
		oldState, err = term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			panic(err)
		}

		go terminalHandler(srv)

		defer term.Restore(int(os.Stdin.Fd()), oldState)
	} else {
		go notRawTerminal(srv)
	}
	srv.Start(timeStart)

	if util.HasArgument("--cpuprof") {
		log.Infoln("Stopped CPU profiler")
		pprof.StopCPUProfile()
	}
	if util.HasArgument("--memprof") {
		log.InfolnClean("Writing memory profile to zeppelin-mem-profile")
		f, _ := os.Create("zeppelin-mem-profile")
		pprof.Lookup("allocs").WriteTo(f, 0)
		f.Close()
	}
}
