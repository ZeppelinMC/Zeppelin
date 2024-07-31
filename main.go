package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"runtime/pprof"
	"strings"
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

	w, err := world.NewWorld("world")
	if err != nil {
		log.Errorlnf("Error loading world: %v", err)
		return
	}

	log.Infof("Binding server to %s:%d\n", cfg.Net.ServerIP, cfg.Net.ServerPort)

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
	} else {
		go notRawTerminal(srv)
	}
	srv.Start(timeStart)

	if rawTerminal {
		term.Restore(int(os.Stdin.Fd()), oldState)
	}
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

func notRawTerminal(srv *server.Server) {
	var line string
	var scanner = bufio.NewScanner(os.Stdin)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

cmdl:
	for {
		select {
		case <-interrupt:
			newText := "\r> stop"
			fmt.Print(newText)
			l := len(line) + 3 // the addtitional 3 chars are "\r> "
			if l > len(newText) {
				fmt.Print(strings.Repeat(" ", l-len(newText)))
			}
			fmt.Println()
			srv.Stop()
			break cmdl
		default:
			if !scanner.Scan() {
				break
			}
			line = scanner.Text()
			srv.CommandManager.Call(line, srv.Console)
			fmt.Print("\r> ")
		}
	}
}

func terminalHandler(srv *server.Server) {
	var char [1]byte
	var currentLine string

charl:
	for {
		os.Stdin.Read(char[:])

		switch char[0] {
		case 8: //backspace
			if currentLine == "" {
				continue
			}
			fmt.Print("\b \b")
			currentLine = currentLine[:len(currentLine)-1]
		case 3: //ctrl-c
			newText := "\r> stop"
			fmt.Print(newText)
			l := len(currentLine) + 3 // the addtitional 3 chars are "\r> "
			if l > len(newText) {
				fmt.Print(strings.Repeat(" ", l-len(newText)))
			}
			fmt.Println()
			srv.Stop()
			break charl
		case 13: // enter
			if currentLine == "" {
				continue
			}
			fmt.Println()
			srv.CommandManager.Call(currentLine, srv.Console)
			currentLine = ""
			fmt.Print("\r> ")
		default:
			char := fmt.Sprintf("%c", char[0])
			currentLine += char
			fmt.Print(char)
		}
	}
}

//8 erase
//3
