package handlers

import (
	"fmt"
	"strings"

	"github.com/dynamitemc/dynamite/server/commands"
)

type controller interface {
	SystemChatMessage(s string) error
}

func ChatCommandPacket(controller controller, graph commands.Graph, content string) {
	args := strings.Split(content, " ")
	cmd := args[0]
	var command *commands.Command
	for _, c := range graph.Commands {
		if c.Name == cmd {
			command = c
		}

		for _, a := range c.Aliases {
			if a == cmd {
				command = c
			}
		}
	}
	if command == nil {
		controller.SystemChatMessage(fmt.Sprintf("§cUnknown or incomplete command, see below for error\n%s§o<--[HERE]", content))
		return
	}
	command.Execute(controller, args)
}
