package core_commands

import (
	"fmt"

	"github.com/aimjel/minecraft/chat"
	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/player"
)

var block_cmd = &commands.Command{
	Name: "block",
	Arguments: []commands.Argument{
		commands.NewVector3Arg("pos"),
	},
	Execute: func(ctx commands.CommandContext) {
		if p, ok := ctx.Executor.(*player.Player); ok {
			x, y, z := p.Position()
			if len(ctx.Arguments) >= 3 {
				x, y, z, _ = ctx.GetVector3("pos")
			}
			b := p.Dimension().Block(int64(x), int64(y), int64(z))
			ctx.Reply(chat.NewMessage(fmt.Sprintf("Block at %d %d %d: %s", int64(x), int64(y), int64(z), b.EncodedName())))
		}
	},
}
