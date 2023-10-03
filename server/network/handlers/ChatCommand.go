package handlers

import (
	"fmt"
	"strings"

	"github.com/dynamitemc/dynamite/server/commands"
)

type controller interface {
	SystemChatMessage(s string) error
	HasPermissions(perms []string) bool
	BroadcastMovement(x1, y1, z1 float64, yaw, pitch float32, ong bool)
	Chat(message string)
	HandleCenterChunk(x1, z1, x2, z2 float64)
}

func ChatCommandPacket(controller controller, graph *commands.Graph, content string) {
	args := strings.Split(content, " ")
	cmd := args[0]
	var command *commands.Command
	for _, c := range graph.Commands {
		if c == nil {
			continue
		}
		if c.Name == cmd {
			command = c
		}

		for _, a := range c.Aliases {
			if a == cmd {
				command = c
			}
		}
	}
	if command == nil || !controller.HasPermissions(command.RequiredPermissions) {
		controller.SystemChatMessage(fmt.Sprintf("§cUnknown or incomplete command, see below for error\n§n%s§r§c§o<--[HERE]", content))
		return
	}
	ctx := commands.CommandContext{
		Arguments:   args[1:],
		Executor:    controller,
		FullCommand: content,
	}
	command.Execute(ctx)
}
