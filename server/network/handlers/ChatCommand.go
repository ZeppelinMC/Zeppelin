package handlers

import (
	"fmt"
	"strings"

	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/player"
)

type controller interface {
	SystemChatMessage(s string) error
	OnGround() bool
	ClientSettings() player.ClientInformation
	Position() (x float64, y float64, z float64)
	Rotation() (yaw float32, pitch float32)
}

func ChatCommandPacket(controller controller, graph *commands.Graph, content string) {
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
		controller.SystemChatMessage(fmt.Sprintf("§cUnknown or incomplete command, see below for error\n§n%s§r§c§o<--[HERE]", content))
		return
	}
	command.Execute(commands.CommandContext{
		Arguments:   args,
		Executor:    controller,
		FullCommand: content,
	})
}
