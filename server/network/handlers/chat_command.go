package handlers

import (
	"fmt"
	"strings"

	"github.com/aimjel/minecraft/chat"
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server/commands"
)

type Controller interface {
	SystemChatMessage(message chat.Message) error
	HasPermissions(perms []string) bool
	BroadcastMovement(id int32, x1, y1, z1 float64, yaw, pitch float32, ong bool, teleport bool)
	Chat(*packet.ChatMessageServer)
	HandleCenterChunk(x1, z1, x2, z2 float64)
	BroadcastPose(pose int32)
	BroadcastSprinting(val bool)
	Hit(entityId int32)
	BroadcastAnimation(animation uint8)
	SendCommandSuggestionsResponse(id int32, start int32, length int32, matches []packet.SuggestionMatch)
	BroadcastSkinData()
	Respawn(dim string)
	BreakBlock(pos uint64)
	Disconnect(reason chat.Message)
	SetClientSettings(p *packet.ClientSettings)
	SetSessionID(id [16]byte, pk, ks []byte, expires int64)
	DropSlot()
	TeleportToEntity(uuid [16]byte)
	UUID() string
	Name() string
	IP() string
	ClearItem(slot int8)
}

func ChatCommandPacket(controller Controller, graph *commands.Graph, log *logger.Logger, content string, timestamp, salt int64, sigs []packet.Argument) {
	log.Info(logger.ParseChat(chat.NewMessage(fmt.Sprintf("[%s] Player %s (%s) issued server command /%s", controller.IP(), controller.Name(), controller.UUID(), content))))
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
		controller.SystemChatMessage(chat.NewMessage(fmt.Sprintf("§cUnknown or incomplete command, see below for error\n§n%s§r§c§o<--[HERE]", content)))
		return
	}
	ctx := commands.CommandContext{
		Command:            command,
		Arguments:          args[1:],
		Executor:           controller,
		FullCommand:        content,
		ArgumentSignatures: sigs,
		Salt:               salt,
		Timestamp:          timestamp,
	}
	command.Execute(ctx)
}
