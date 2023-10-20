package core_commands

import (
	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/server/commands"
)

func getServer(executor interface{}) *server.Server {
	if p, ok := executor.(*server.Session); ok {
		return p.Server
	} else if c, ok := executor.(*server.ConsoleExecutor); ok {
		return c.Server
	}
	return nil
}

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
