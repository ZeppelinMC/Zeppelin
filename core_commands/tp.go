package core_commands

import (
	"fmt"
	"strconv"

	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/server/commands"
)

var tp_cmd = &commands.Command{
	Name:                "tp",
	RequiredPermissions: []string{"server.command.op"},
	Aliases:             []string{"teleport"},
	Arguments: []commands.Argument{
		commands.NewEntityArgument("targets", commands.EntityPlayerOnly),
		commands.NewEntityArgument("destination", commands.EntityPlayerOnly).SetAlternative(commands.NewVector3Argument("location")),
	},
	Execute: func(ctx commands.CommandContext) {
		srv := getServer(ctx.Executor)
		switch len(ctx.Arguments) {
		case 1:
			{
				if exe, ok := ctx.Executor.(*server.Session); !ok {
					ctx.Incomplete()
					return
				} else {
					player := srv.FindPlayer(ctx.Arguments[0])
					x, y, z := player.Position()
					yaw, pitch := exe.Rotation()
					exe.Teleport(x, y, z, yaw, pitch)
					ctx.Reply(fmt.Sprintf("Teleported %s to %s", exe.Name(), player.Name()))
				}
			}
		case 2:
			{
				// Teleport player to player
				player1 := srv.FindPlayer(ctx.Arguments[0])
				player2 := srv.FindPlayer(ctx.Arguments[1])
				x, y, z := player2.Position()
				yaw, pitch := player1.Rotation()
				player1.Teleport(x, y, z, yaw, pitch)

				ctx.Reply(fmt.Sprintf("Teleported %s to %s", player1.Name(), player2.Name()))
			}
		case 3:
			{
				// Teleport executor to coordinates
				if exe, ok := ctx.Executor.(*server.Session); !ok {
					ctx.Incomplete()
				} else {
					x, err := strconv.ParseFloat(ctx.Arguments[0], 64)
					if err != nil {
						ctx.Error("Invalid x position")
						return
					}
					y, err := strconv.ParseFloat(ctx.Arguments[1], 64)
					if err != nil {
						ctx.Error("Invalid y position")
						return
					}
					z, err := strconv.ParseFloat(ctx.Arguments[2], 64)
					if err != nil {
						ctx.Error("Invalid x position")
						return
					}
					yaw, pitch := exe.Rotation()

					exe.Teleport(x, y, z, yaw, pitch)
				}
			}
		case 4:
			{
				// teleport player to coordinates
				player := srv.FindPlayer(ctx.Arguments[0])
				x, err := strconv.ParseFloat(ctx.Arguments[1], 64)
				if err != nil {
					ctx.Error("Invalid x position")
					return
				}
				y, err := strconv.ParseFloat(ctx.Arguments[2], 64)
				if err != nil {
					ctx.Error("Invalid y position")
					return
				}
				z, err := strconv.ParseFloat(ctx.Arguments[3], 64)
				if err != nil {
					ctx.Error("Invalid x position")
					return
				}

				yaw, pitch := player.Rotation()
				player.Teleport(x, y, z, yaw, pitch)

				ctx.Reply(fmt.Sprintf("Teleported %s", player.Name()))
			}
		default:
			ctx.Incomplete()
		}
	},
}
