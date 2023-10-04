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
	fmt.Print("> ")

	reader := bufio.NewReader(os.Stdin)
	content, _ := reader.ReadString('\n')
	content = strings.TrimSpace(content)

	args := strings.Split(content, " ")
	cmd := args[0]

	if cmd == "" {
		return
	}

	defer srv.ScanConsole()
	command := srv.CommandGraph.FindCommand(cmd)
	if command == nil {
		fmt.Println(commands.ParseChat(fmt.Sprintf("§cUnknown or incomplete command, see below for error\n§n%s§r§c§o<--[HERE]", cmd)))
		return
	}
	command.Execute(commands.CommandContext{
		Arguments:   args[1:],
		Executor:    &ConsoleExecutor{Server: srv},
		FullCommand: content,
	})
}
