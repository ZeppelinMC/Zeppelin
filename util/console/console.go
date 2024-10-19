package console

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"unsafe"

	"github.com/zeppelinmc/zeppelin/server"
)

func GetFlag(name string) (string, bool) {
	name = "--" + name + "="
	for _, a := range os.Args {
		if i := strings.Index(a, name); i == 0 {
			if len(name)+i < len(a) {
				return a[len(name)+i:], true
			}
		}
	}

	return "", false
}

func StartConsole(srv *server.Server) {
	var line string
	var scanner = bufio.NewScanner(os.Stdin)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill)

	go func() {
		for range interrupt {
			newText := "\r> stop"
			fmt.Print(newText)
			l := len(line) + 3 // the addtitional 3 chars are "\r> "
			if l > len(newText) {
				fmt.Print(strings.Repeat(" ", l-len(newText)))
			}
			fmt.Println()
			srv.Stop()
			os.Stdin.Close()
		}
	}()

	for {
		if !scanner.Scan() {
			break
		}
		line = scanner.Text()
		srv.CommandManager.Call(line, srv.Console)
		fmt.Print("\r> ")
	}
}

func StartRawConsole(srv *server.Server) {
	var char [1]byte
	var currentLine string

	var previousLines []string
	var previousLinesIndex int

	var currentLineIndex int

charl:
	for {
		os.Stdin.Read(char[:])
		if char[0] == 27 {
			os.Stdin.Read(char[:])
			if char[0] != 91 {
				continue
			}
			os.Stdin.Read(char[:])

			switch char[0] {
			case 'A': //up
				if len(previousLines[previousLinesIndex:]) == 0 {
					continue
				}
				l := previousLines[previousLinesIndex]
				previousLinesIndex++

				fmt.Print("\r> ", l)
				if l, l0 := len(currentLine), len(l); l > l0 {
					fmt.Print(strings.Repeat(" ", l-l0))
				}
				currentLineIndex = len(l) - 1
				currentLine = l
			case 'B': //down
			case 'C': //right
				if currentLineIndex == len(currentLine) {
					continue
				}
				fmt.Printf("%c", currentLine[currentLineIndex])
				currentLineIndex++
			case 'D': //left
				if currentLineIndex == 0 {
					continue
				}
				currentLineIndex--
				fmt.Print("\b")
			}
			continue
		}
		switch char[0] {
		case '\b', 127: //backspace
			if len(currentLine) == 0 {
				continue
			}
			fmt.Print("\b \b")
			currentLine = currentLine[:len(currentLine)-1]
			currentLineIndex--
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
			previousLines = append([]string{currentLine}, previousLines...)
			currentLine = ""
			currentLineIndex = 0
			fmt.Print("\r> ")
		default:
			if currentLineIndex < len(currentLine) {
				*unsafe.StringData(currentLine[currentLineIndex:]) = char[0]
			} else {
				currentLine += fmt.Sprintf("%c", char[0])
			}
			currentLineIndex++
			fmt.Printf("%c", char[0])
		}
	}
}

//8 erase
//3
