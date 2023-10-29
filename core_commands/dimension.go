package core_commands

import (
	"github.com/aimjel/minecraft/chat"
	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/server/commands"
)

var dimension_cmd = &commands.Command{
	Name:                "dimension",
	RequiredPermissions: []string{"server.command.dimension"},
	Arguments: []commands.Argument{
		commands.NewDimensionArgument("dimension"),
	},
	Execute: func(ctx commands.CommandContext) {
		if len(ctx.Arguments) == 0 {
			ctx.Incomplete()
			return
		}
		if p, ok := ctx.Executor.(*server.Session); ok {
			p.Respawn(ctx.Arguments[0])
			ctx.Reply(chat.NewMessage("Switched dimension to " + ctx.Arguments[0]))
		} else {
			ctx.Error("This command can only be used by players")
			return
		}
	},
}
