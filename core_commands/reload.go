package core_commands

import (
	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/server/commands"
)

var reload_cmd = &commands.Command{
	Name:                "reload",
	Aliases:             []string{"rl"},
	RequiredPermissions: []string{"server.reload"},
	Execute: func(ctx commands.CommandContext) {
		srv := ctx.Executor.(*server.PlayerController).Server
		srv.Reload()
		ctx.Reply(srv.Config.Messages.ReloadComplete)
	},
}
