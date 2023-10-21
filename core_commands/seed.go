package core_commands

import (
	"fmt"

	"github.com/dynamitemc/dynamite/server/commands"
)

var seed_cmd = &commands.Command{
	Name:                "seed",
	RequiredPermissions: []string{"server.command.seed"},
	Execute: func(ctx commands.CommandContext) {
		server := getServer(ctx.Executor)
		ctx.Reply(fmt.Sprint(server.World.Seed()))
	},
}
