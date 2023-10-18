package core_commands

import (
	"github.com/dynamitemc/dynamite/server/commands"
	"math"
)

var low, high uint64 = math.Float64bits(1), math.Float64bits(2)
var test_cmd = &commands.Command{
	Name: "test",
	Arguments: []commands.Argument{
		commands.NewFloatArg("wqeqwf").MinMax(low, high),
	},
	Execute: func(ctx commands.CommandContext) {
		ctx.Reply("hi")
	},
}
