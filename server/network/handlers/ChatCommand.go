package handlers

import (
	"strings"

	"github.com/dynamitemc/dynamite/server/commands"
)

func ChatCommandPacket(controller interface{}, graph commands.Graph, content string) {
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
		return
	}
	command.Execute(controller, args)
}
