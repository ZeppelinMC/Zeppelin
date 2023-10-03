package core_commands

import (
	"github.com/dynamitemc/dynamite/server/commands"
)

var restart_cmd = &commands.Command{
	Name:                "restart",
	Aliases:             []string{"rs"},
	RequiredPermissions: []string{"server.command.restart"},
	Execute: func(ctx commands.CommandContext) {
		panic("restarted")
	},
}
