package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/zeppelinmc/zeppelin/server"
)

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
		case '\b', 127: //backspace
			if len(currentLine) == 0 {
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
