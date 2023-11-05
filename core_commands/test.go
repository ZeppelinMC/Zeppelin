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
		commands.NewStrArg("d", commands.SingleWord),
		commands.NewBoolArg("gay"),
		commands.NewFloatArg("f"),
		commands.NewVector3Argument("pos"),
	},
	Execute: func(ctx commands.CommandContext) {
		d, _ := ctx.GetString("d")
		gay, _ := ctx.GetBool("gay")
		f, _ := ctx.GetFloat64("f")
		x, y, z, _ := ctx.GetVector3("pos")

		ctx.Reply(chat.NewMessage(fmt.Sprintf("%s %v %f    %f %f %f", d, gay, f, x, y, z)))
	},
}
