package core_commands

import (
	"strings"

	"github.com/dynamitemc/dynamite/server/commands"
)

var ban_cmd = &commands.Command{
	Name:                "ban",
	RequiredPermissions: []string{"server.command.ban"},
	Arguments: []commands.Argument{
		commands.NewEntityArg("player", commands.EntityPlayerOnly),
		commands.NewStrArg("reason", commands.GreedyPhrase),
	},
	Execute: func(ctx commands.CommandContext) {
		if len(ctx.Arguments) == 0 {
			ctx.Incomplete()
			return
		}
		server := getServer(ctx.Executor)
		playerName := ctx.Arguments[0]
		reason := server.Translate("disconnect.banned", nil)
		if len(ctx.Arguments) > 1 {
			reason = server.Translate("disconnect.banned.reason", map[string]string{"reason": strings.Join(ctx.Arguments[1:], " ")})
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
