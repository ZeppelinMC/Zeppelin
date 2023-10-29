package core_commands

import (
	"strings"

	"github.com/aimjel/minecraft/chat"
	"github.com/dynamitemc/dynamite/server/commands"
)

var ban_cmd = &commands.Command{
	Name:                "ban",
	RequiredPermissions: []string{"server.command.ban"},
	Arguments: []commands.Argument{
		commands.NewEntityArgument("player", commands.EntityPlayerOnly),
		commands.NewStrArg("reason", commands.GreedyPhrase),
	},
	Execute: func(ctx commands.CommandContext) {
		if len(ctx.Arguments) == 0 {
			ctx.Incomplete()
			return
		}
		server := getServer(ctx.Executor)
		playerName := ctx.Arguments[0]
		reason := server.Translate("multiplayer.disconnect.banned")
		if len(ctx.Arguments) > 1 {
			reason = server.Translate("multiplayer.disconnect.banned.reason", chat.NewMessage(strings.Join(ctx.Arguments[1:], " ")))
		}
		player := server.FindPlayer(playerName)
		if player == nil {
			ctx.Error("No player was found")
			return
		}
		server.Ban(player, strings.Join(ctx.Arguments[1:], " "))
		player.Disconnect(reason)
	},
}
