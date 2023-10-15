package core_commands

import (
	"strconv"

	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/server/commands"
)

var tp_cmd = &commands.Command{
	Name:                "tp",
	Aliases:             []string{"teleport"},
	RequiredPermissions: []string{"server.command.tp"},
	Arguments: []commands.Argument{
		commands.NewEntityArgument("player", commands.EntityPlayerOnly),
		commands.NewFloatArgument("x", struct {
			Min *uint64
			Max *uint64
		}{Min: &m, Max: &ma}),
		commands.NewFloatArgument("y", struct {
			Min *uint64
			Max *uint64
		}{Min: &m, Max: &ma}),
		commands.NewFloatArgument("z", struct {
			Min *uint64
			Max *uint64
		}{Min: &m, Max: &ma}),
	},
	Execute: func(ctx commands.CommandContext) {
		srv := getServer(ctx.Executor)

		if len(ctx.Arguments) == 0 {
			ctx.Incomplete()
			return
		}

		switch len(ctx.Arguments) {
		case 1:
			{
				if p, ok := ctx.Executor.(*server.PlayerController); ok {
					// from executor, to a player
					target := srv.FindPlayer(ctx.Arguments[0])
					if target == nil {
						ctx.Reply("Player not found")
						return
					}
					targetX, targetY, targetZ := target.Position()
					targetYaw, targetPitch := p.Rotation() // do not change rotation

					p.Teleport(targetX, targetY, targetZ, targetYaw, targetPitch)
					ctx.Reply("Teleported " + p.Name())
				} else {
					ctx.Incomplete()
					return
				}
			}
		case 2:
			{
				// from player to player
				// can be executed by both console and player

				source := srv.FindPlayer(ctx.Arguments[0])
				if source == nil {
					ctx.Reply("Player not found")
					return
				}

				target := srv.FindPlayer(ctx.Arguments[1])
				if target == nil {
					ctx.Reply("Player not found")
					return
				}

				targetX, targetY, targetZ := target.Position()
				targetYaw, targetPitch := source.Rotation() // do not change rotation

				source.Teleport(targetX, targetY, targetZ, targetYaw, targetPitch)
				ctx.Reply("Teleported " + source.Name() + " to " + target.Name())
			}

		case 3:
			{
				// teleport executor to coordinates

				if p, ok := ctx.Executor.(*server.PlayerController); ok {
					x, _ := strconv.ParseFloat(ctx.Arguments[0], 64)
					y, _ := strconv.ParseFloat(ctx.Arguments[1], 64)
					z, _ := strconv.ParseFloat(ctx.Arguments[2], 64)

					targetYaw, targetPitch := p.Rotation() // do not change rotation

					p.Teleport(x, y, z, targetYaw, targetPitch)
					ctx.Reply("Teleported " + p.Name())
				} else {
					ctx.Incomplete()
					return
				}
			}
		case 4:
			{
				// teleport player to coordinates
				// can be executed by both console and player

				target := srv.FindPlayer(ctx.Arguments[0])
				if target == nil {
					ctx.Reply("Player not found")
					return
				}

				x, _ := strconv.ParseFloat(ctx.Arguments[1], 64)
				y, _ := strconv.ParseFloat(ctx.Arguments[2], 64)
				z, _ := strconv.ParseFloat(ctx.Arguments[3], 64)

				targetYaw, targetPitch := target.Rotation() // do not change rotation

				target.Teleport(x, y, z, targetYaw, targetPitch)
				ctx.Reply("Teleported " + target.Name())
			}
		}
	},
}
