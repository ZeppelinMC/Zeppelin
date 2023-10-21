package core_commands

import (
	"github.com/dynamitemc/dynamite/server/commands"
)

var reload_cmd = &commands.Command{
	Name:                "reload",
	Aliases:             []string{"rl"},
	RequiredPermissions: []string{"server.command.reload"},
	Execute: func(ctx commands.CommandContext) {
		srv := getServer(ctx.Executor)
		srv.Reload()
		ctx.Reply(srv.Config.Messages.ReloadComplete)
	},
}
