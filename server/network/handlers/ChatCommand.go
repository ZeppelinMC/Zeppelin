package handlers

import (
	"fmt"
	"strings"

	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/world"
)

type Controller interface {
	SystemChatMessage(s string) error
	HasPermissions(perms []string) bool
	BroadcastMovement(id int32, x1, y1, z1 float64, yaw, pitch float32, ong bool, teleport bool)
	Chat(message string)
	HandleCenterChunk(x1, z1, x2, z2 float64)
	BroadcastPose(pose int32)
	BroadcastSprinting(val bool)
	Hit(entityId int32)
	BroadcastAnimation(animation uint8)
	SendCommandSuggestionsResponse(id int32, start int32, length int32, matches []packet.SuggestionMatch)
	BroadcastSkinData()
	Respawn(d *world.Dimension)
	PlaylistUpdate()
}

func ChatCommandPacket(controller Controller, graph *commands.Graph, content string) {
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
