package core_commands

import (
	"github.com/aimjel/minecraft/chat"
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
		var player *server.Session
		if len(ctx.Arguments) == 0 {
			if p, ok := ctx.Executor.(*server.Session); !ok {
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
		ctx.Reply(player.Server.Translate("commands.kill.success.single", chat.NewMessage(name)))
		player.Server.GlobalMessage(chat.NewMessage(name + " was killed"))
	},
}
