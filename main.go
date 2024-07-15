package main

import (
	"fmt"
	"os"
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

func main() {
	log.Infoln("Aether 1.21 Minecraft server")
	log.Infof("Running on platform %s-%s\n", runtime.GOOS, runtime.GOARCH)
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

	log.Infof("Binding server to %s:%d TCP\n", cfg.ServerIP, cfg.ServerPort)
	srv, err := cfg.New()
	if err != nil {
		log.Errorln("Error binding server:", err)
		return
	}
	srv.Start(timeStart, terminalHandler)
}

func terminalHandler(srv *server.Server) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	fmt.Printf("> ")

	var char [1]byte
	var currentLine string
	for {
		os.Stdin.Read(char[:])

		switch char[0] {
		case 8:
			if currentLine == "" {
				continue
			}
			fmt.Print("\b \b")
			currentLine = currentLine[:len(currentLine)-1]
		case 3:
			newText := "\r> stop"
			fmt.Print(newText)
			l := len(currentLine) + 3 // the addtitional 3 chars are "\r> "
			if l > len(newText) {
				fmt.Print(strings.Repeat(" ", l-len(newText)))
			}
			fmt.Println()
			srv.Stop()
		case 13:
			if currentLine == "" {
				continue
			}
			currentLine = ""
			fmt.Print("\n> ")
		default:
			char := fmt.Sprintf("%c", char[0])
			currentLine += char
			fmt.Print(char)
		}
	}
}

//8 erase
//3
