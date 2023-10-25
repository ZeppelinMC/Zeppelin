package core_commands

import (
	"fmt"

	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/server/commands"
)

var list_cmd = &commands.Command{
	Name: "list",
	Execute: func(ctx commands.CommandContext) {
		switch ex := ctx.Executor.(type) {

		case *server.Session:
			ctx.Reply(fmt.Sprintf("%v players online", ex.Server.PlayerCount()))

		case *server.ConsoleExecutor:
			ctx.Reply(fmt.Sprintf("%v players online", ex.Server.PlayerCount()))
		}
	},
}
