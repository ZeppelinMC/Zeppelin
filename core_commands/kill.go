package core_commands

import (
	"fmt"

	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/server/commands"
)

var kill_cmd = &commands.Command{
	Name:                "kill",
	RequiredPermissions: []string{"server.command.kill"},
	Arguments: []commands.Argument{
		commands.NewEntityArgument("player", commands.EntityPlayerOnly),
	},
	Execute: func(ctx commands.CommandContext) {
		var player *server.PlayerController
		if len(ctx.Arguments) == 0 {
			if p, ok := ctx.Executor.(*server.PlayerController); !ok {
				ctx.Incomplete()
				return
			} else {
				player = p
			}
		} else {
			p := getServer(ctx.Executor).FindPlayer(ctx.Arguments[0])
			if p == nil {
				ctx.Error("No player was found")
				return
			}
			player = p
		}
		name := player.Name()
		player.Kill(name + " was killed")
		ctx.Reply(fmt.Sprintf("Killed %s", name))
		player.Server.GlobalMessage(name+"was killed", nil)
	},
}
