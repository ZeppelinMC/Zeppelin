package core_commands

import (
	"github.com/dynamitemc/dynamite/server/commands"
)

var unban_cmd = &commands.Command{
	Name:                "unban",
	Aliases:             []string{"pardon"},
	RequiredPermissions: []string{"server.command.unban"},
	Arguments: []commands.Argument{
		commands.NewStringArgument("player", commands.StringSingleWord),
	},
	Execute: func(ctx commands.CommandContext) {
		if len(ctx.Arguments) == 0 {
			ctx.Incomplete()
			return
		}
		playerName := ctx.Arguments[0]
		server := getServer(ctx.Executor)

		server.Unban(playerName)
		ctx.Reply("Unbanned " + playerName)
	},
}
