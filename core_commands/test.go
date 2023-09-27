package core_commands

import (
	"github.com/dynamitemc/dynamite/server/commands"
)

var test_cmd = &commands.Command{
	Name: "test",
	Execute: func(ctx commands.CommandContext) {
		ctx.Reply("hi")
	},
}
