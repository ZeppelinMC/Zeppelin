package core_commands

import (
	"strconv"

	"github.com/dynamitemc/dynamite/server/commands"
)

var summon_cmd = &commands.Command{
	Name:                "summon",
	Aliases:             []string{},
	RequiredPermissions: []string{"server.command.summon"},
	Arguments: []commands.Argument{
		commands.NewResourceKeyArg("entity", "minecraft:entity_type"),
		commands.NewVector3Arg("location"),
	},
	Execute: func(ctx commands.CommandContext) {
		if len(ctx.Arguments) < 4 {
			ctx.Incomplete()
			return
		}
		typ := ctx.Arguments[0]
		x1, y1, z1 := ctx.Arguments[1], ctx.Arguments[2], ctx.Arguments[3]
		x, err := strconv.ParseFloat(x1, 64)
		if err != nil {
			ctx.Error("Invalid x position")
			return
		}
		y, err := strconv.ParseFloat(y1, 64)
		if err != nil {
			ctx.Error("Invalid y position")
			return
		}
		z, err := strconv.ParseFloat(z1, 64)
		if err != nil {
			ctx.Error("Invalid z position")
			return
		}
		srv := getServer(ctx.Executor)
		srv.SpawnEntity(typ, x, y, z)
	},
}
