package core_commands

import (
	"fmt"
	"strings"

	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/player"
)

var msg_cmd = &commands.Command{
	Name:                "msg",
	RequiredPermissions: []string{"server.chat"},
	Arguments: []commands.Argument{
		commands.NewEntityArg("targets", commands.EntitySingle),
		commands.NewStrArg("message", commands.GreedyPhrase),
	},
	Execute: func(ctx commands.CommandContext) {
		var pl *player.Player
		if len(ctx.Arguments) < 2 {
			ctx.Incomplete()
			return
		}
		if p, ok := ctx.Executor.(*player.Player); !ok {
			//todo implement
			return
		} else {
			pl = getServer(ctx.Executor).FindPlayer(ctx.Arguments[0])
			if pl == nil {
				ctx.Error("No player was found")
				return
			}
			fmt.Println(ctx.ArgumentSignatures)
			p.Whisper(pl, strings.Join(ctx.Arguments[1:], " "), ctx.Timestamp, ctx.Salt, nil) //, ctx.ArgumentSignatures[1].Signature)
		}
	},
}
