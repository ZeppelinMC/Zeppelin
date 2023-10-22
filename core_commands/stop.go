package core_commands

import (
	"github.com/dynamitemc/dynamite/server/commands"
)

var stop_cmd = &commands.Command{
	Name:                "stop",
	Aliases:             []string{},
	RequiredPermissions: []string{"server.command.stop"},
	Execute: func(ctx commands.CommandContext) {
		srv := getServer(ctx.Executor)
		srv.Close()
	},
}
