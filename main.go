package main

import (
	_ "embed"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"slices"
	"strconv"
	"time"

	"github.com/zeppelinmc/zeppelin/commands"
	"github.com/zeppelinmc/zeppelin/protocol/properties"
	"github.com/zeppelinmc/zeppelin/server"
	"github.com/zeppelinmc/zeppelin/server/command"
	"github.com/zeppelinmc/zeppelin/server/world"
	"github.com/zeppelinmc/zeppelin/server/world/chunk/section"
	"github.com/zeppelinmc/zeppelin/util"
	"github.com/zeppelinmc/zeppelin/util/console"
	"github.com/zeppelinmc/zeppelin/util/log"
	"golang.org/x/term"
)

var timeStart = time.Now()

func main() {
	log.Infolnf("Zeppelin 1.21 Minecraft server with %s on platform %s-%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	if !loadStates() {
		return
	}

	max, ok := console.GetFlag("xmem")
	if ok {
		m, err := util.ParseSizeUnit(max)
		if err == nil {
			debug.SetMemoryLimit(int64(m))
			log.Infolnf("Memory usage is limited to %s", util.FormatSizeUnit(m))
		}
	}

	if slices.Index(os.Args, "--cpuprof") != -1 {
		if f, err := os.Create("zeppelin-cpu-profile"); err == nil {
			pprof.StartCPUProfile(f)
			log.Infoln("Started CPU profiler (writing to zeppelin-cpu-profile)")

			defer func() {
				log.Infoln("Stopped CPU profiler")
				pprof.StopCPUProfile()
				f.Close()
			}()
		}
	}

	if slices.Index(os.Args, "--memprof") != -1 {
		defer func() {
			log.InfolnClean("Writing memory profile to zeppelin-mem-profile")
			f, _ := os.Create("zeppelin-mem-profile")
			pprof.Lookup("allocs").WriteTo(f, 0)
			f.Close()
		}()
	}

	cfg := loadConfig()
	if cfg.LevelSeed == "" {
		cfg.LevelSeed = strconv.FormatInt(rand.Int63(), 10)
	}

	w, err := world.NewWorld(cfg)
	if err != nil {
		log.Errorlnf("Error preparing level: %v", w)
		return
	}

	log.Infof("Binding server to %s:%d\n", cfg.ServerIp, cfg.ServerPort)

	rawTerminal := slices.Index(os.Args, "--no-raw-terminal") == -1

	srv, err := server.New(cfg, w)
	if err != nil {
		log.Errorln("Error binding server:", err)
		return
	}
	srv.CommandManager = command.NewManager(srv, commands.Commands...)
	if rawTerminal {
		oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			panic(err)
		}

		go console.StartRawConsole(srv)

		defer term.Restore(int(os.Stdin.Fd()), oldState)
	} else {
		go console.StartConsole(srv)
	}
	srv.Start(timeStart)
}

func loadConfig() properties.ServerProperties {
	file, err := os.ReadFile("server.properties")
	if err != nil {
		file, err := os.Create("server.properties")
		if err == nil {
			properties.Marshal(file, properties.Default)
			file.Close()
		}
		return properties.Default
	}
	var cfg properties.ServerProperties

	err = properties.Unmarshal(string(file), &cfg)
	if err != nil {
		cfg = properties.Default
	}

	return cfg
}

func loadStates() (ok bool) {
	os.Mkdir("resources", 0755)
	f, err := os.OpenFile("./resources/blocks", os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		log.ErrorlnClean("Error opening block states file: ", err)
		return false
	}

	if err := section.ImportStates(f); err == nil {
		return true
	}

	log.Info("Downloading block states file")

	res, err := http.Get(statesfilei)
	if err != nil {
		fmt.Println("... error")
		log.ErrorlnClean("Error downloading block states file: ", err)
		return false
	}
	if _, err := f.ReadFrom(res.Body); err != nil {
		fmt.Println("... error")
		log.ErrorlnClean("Error downloading block states file: ", err)
		return false
	}

	f.Seek(0, 0)

	fmt.Println("... done")

	if err := section.ImportStates(f); err != nil {
		log.ErrorlnClean("An error occured while loading block states file: ", err)

		return false
	}

	return true
}

var statesfilei = "https://github.com/ZeppelinMC/resources/raw/main/blocks"
