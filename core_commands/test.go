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
	},
	Execute: func(ctx commands.CommandContext) {
		d, _ := ctx.GetString("d")
		gay, _ := ctx.GetBool("gay")
		f, _ := ctx.GetFloat64("f")

		ctx.Reply(chat.NewMessage(fmt.Sprintf("%s %v %f", d, gay, f)))
	},
}
