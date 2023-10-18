package core_commands

import (
	"strings"

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
		playerName := ctx.Arguments[0]
		reason := "Banned by an operator"
		if len(ctx.Arguments) > 1 {
			reason = strings.Join(ctx.Arguments[1:], " ")
		}
		server := getServer(ctx.Executor)
		player := server.FindPlayer(playerName)
		if player == nil {
			ctx.Error("No player was found")
			return
		}
		server.Ban(player, reason)
		player.Disconnect(reason)
	},
}
