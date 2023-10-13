package server

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/dynamitemc/dynamite/server/commands"
)

type ConsoleExecutor struct {
	Server *Server
}

func (srv *Server) ScanConsole() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			continue
		}

		content := strings.TrimSpace(txt)
		args := strings.Split(content, " ")

		command := srv.commandGraph.FindCommand(args[0])
		if command == nil {
			srv.Logger.Print(fmt.Sprintf("&cUnknown or incomplete command, see below for error\n&n%s&r&c&o<--[HERE]", args[0]))
			return
		}
		command.Execute(commands.CommandContext{
			Arguments:   args[1:],
			Executor:    &ConsoleExecutor{Server: srv},
			FullCommand: content,
		})
	}

	if err := scanner.Err(); err != nil {
		srv.Logger.Error("%v scanning console", err)
	}
}
