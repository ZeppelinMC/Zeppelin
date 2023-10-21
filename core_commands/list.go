package core_commands

import (
	"fmt"

	"github.com/dynamitemc/dynamite/server/commands"
)

var list_cmd = &commands.Command{
	Name:                "list",
	RequiredPermissions: []string{"server.command.list"},
	Arguments: []commands.Argument{
		commands.NewStrArg("uuids", commands.SingleWord), // should suggest the uuids string but it doesn't
	},
	Execute: func(ctx commands.CommandContext) {
		srv := getServer(ctx.Executor)
		players := srv.Players
		if len(players) == 0 {
			ctx.Reply("No players online")
			return
		}
		msg := fmt.Sprintf("There are %d of a max of %d players online:", len(players), srv.Config.MaxPlayers)
		for _, p := range players {
			if len(ctx.Arguments) == 1 && ctx.Arguments[0] == "uuids" {
				msg += fmt.Sprintf("\n(%s)", p.UUID)
			} else {
				msg += fmt.Sprintf("\n - %s", p.Name())
			}
		}
		ctx.Reply(msg)
	},
}
