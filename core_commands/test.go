package core_commands

import (
	"math"

	"github.com/dynamitemc/dynamite/server/commands"
)

var m, ma uint64 = math.Float64bits(1), math.Float64bits(2)
var test_cmd = &commands.Command{
	Name: "test",
	Arguments: []commands.Argument{
		commands.NewFloatArgument("wqeqwf", struct {
			Min *uint64
			Max *uint64
		}{Min: &m, Max: &ma}),
	},
	Execute: func(ctx commands.CommandContext) {
		ctx.Reply("hi")
	},
}
