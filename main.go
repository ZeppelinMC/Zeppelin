package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"

	"github.com/dynamitemc/aether/log"
	"github.com/dynamitemc/aether/net/registry"
	"github.com/dynamitemc/aether/server"
	"github.com/dynamitemc/aether/server/world/region/blocks"
	"golang.org/x/term"
)

var timeStart = time.Now()

func hasArgument(name string) bool {
	for _, arg := range os.Args {
		if arg == name {
			return true
		}
	}
	return false
}

func main() {
	log.Infoln("Aether 1.21 Minecraft server")
	log.Infof("Running on %s on platform %s-%s\n", runtime.Version(), runtime.GOOS, runtime.GOARCH)
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

	log.Infof("Binding server to %s:%d TCP\n", cfg.Net.ServerIP, cfg.Net.ServerPort)

	rawTerminal := !hasArgument("--no-raw-terminal")

	srv, err := cfg.New()
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
