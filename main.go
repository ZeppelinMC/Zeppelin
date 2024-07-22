package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"

	"github.com/zeppelinmc/zeppelin/log"
	"github.com/zeppelinmc/zeppelin/server"
	"github.com/zeppelinmc/zeppelin/server/world/region/blocks"
	"github.com/zeppelinmc/zeppelin/util"
	"golang.org/x/term"
)

var timeStart = time.Now()

func main() {
	log.Infolnf("Zeppelin 1.21 Minecraft server with %s on platform %s-%s", runtime.Version(), runtime.GOOS, runtime.GOARCH)
	if err := blocks.LoadBlockCache(); err != nil {
		log.Errorln("Error loading server registries:", err)
		return
	}

	cfg := loadConfig()

	log.Infof("Binding server to %s:%d\n", cfg.Net.ServerIP, cfg.Net.ServerPort)

	rawTerminal := !util.HasArgument("--no-raw-terminal")

	srv, err := server.New(cfg)
	if err != nil {
		log.Errorln("Error binding server:", err)
		return
	}
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
}

func notRawTerminal(srv *server.Server) {
	var line string

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
			fmt.Scanln(&line)
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
			currentLine = ""
			fmt.Print("\n\r> ")
		default:
			char := fmt.Sprintf("%c", char[0])
			currentLine += char
			fmt.Print(char)
		}
	}
}

//8 erase
//3
