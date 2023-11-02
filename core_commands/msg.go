package core_commands

import (
	"fmt"
	"strings"

	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/server/commands"
)

var msg_cmd = &commands.Command{
	Name:                "msg",
	RequiredPermissions: []string{"server.chat"},
	Arguments: []commands.Argument{
		commands.NewEntityArgument("targets", commands.EntitySingle),
		commands.NewStrArg("message", commands.GreedyPhrase),
	},
	Execute: func(ctx commands.CommandContext) {
		var player *server.Session
		if len(ctx.Arguments) < 2 {
			ctx.Incomplete()
			return
		}
		if p, ok := ctx.Executor.(*server.Session); !ok {
			//todo implement
			return
		} else {
			player = getServer(ctx.Executor).FindPlayer(ctx.Arguments[0])
			if player == nil {
				ctx.Error("No player was found")
				return
			}
			fmt.Println(ctx.ArgumentSignatures)
			p.Whisper(player, strings.Join(ctx.Arguments[1:], " "), ctx.Timestamp, ctx.Salt, nil) //, ctx.ArgumentSignatures[1].Signature)
		}
	},
}
