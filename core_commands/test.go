package core_commands

import (
	"fmt"

	"github.com/aimjel/minecraft/chat"
	"github.com/dynamitemc/dynamite/server/commands"
)

var test_cmd = &commands.Command{
	Name:                "test",
	RequiredPermissions: []string{"server.command.test"},
	Arguments: []commands.Argument{
		commands.NewVector3Arg("pos1"),
		commands.NewStrArg("d", commands.SingleWord),
		commands.NewBoolArg("gay"),
		commands.NewIntArg("f"),
		commands.NewVector3Arg("pos"),
	},
	Execute: func(ctx commands.CommandContext) {
		x, y, z, _ := ctx.GetVector3("pos1")
		d, _ := ctx.GetString("d")
		gay, _ := ctx.GetBool("gay")
		f, _ := ctx.GetInt64("f")
		x1, y1, z1, _ := ctx.GetVector3("pos")

		ctx.Reply(chat.NewMessage(fmt.Sprintf("%f %f %f %s %v %d %f %f %f", x, y, z, d, gay, f, x1, y1, z1)))
	},
}
