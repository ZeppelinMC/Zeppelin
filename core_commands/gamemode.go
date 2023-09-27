package core_commands

import (
	"fmt"
	"strings"

	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/server/commands"
	p "github.com/dynamitemc/dynamite/server/player"
)

func pascalify(str string) (res string) {
	sp := strings.Split(str, " ")
	for _, w := range sp {
		spp := strings.Split(w, "")
		res += strings.ToUpper(spp[0]) + strings.ToLower(strings.Join(spp[1:], ""))
	}
	return
}

var gamemode_cmd = &commands.Command{
	Name:                "gamemode",
	RequiredPermissions: []string{"server.command.gamemode"},
	Arguments: []commands.Argument{
		{
			Name:     "mode",
			ParserID: 39,
		},
		{
			Name:     "player",
			ParserID: 6,
			Properties: commands.Properties{
				Flags: 0x02,
			},
		},
	},
	Execute: func(ctx commands.CommandContext) {
		if len(ctx.Arguments) == 0 {
			ctx.Incomplete()
			return
		}
		gm := p.Gamemode(ctx.Arguments[0])
		if gm == -1 {
			ctx.ErrorAt(fmt.Sprintf("Unknown game mode: %s", ctx.Arguments[0]))
			return
		}
		var player *server.PlayerController
		if len(ctx.Arguments) == 1 {
			if ctx.IsConsole {
				ctx.Incomplete()
				return
			}
			player = ctx.Executor.(*server.PlayerController)
		} else {
			p := ctx.Executor.(*server.PlayerController).Server.FindPlayer(ctx.Arguments[1])
			if p == nil {
				ctx.Error("No player was found")
				return
			}
			player = p
		}
		if player.GameMode() == byte(gm) {
			return
		}
		player.SetGameMode(byte(gm))
		msg := fmt.Sprintf("Set %s's game mode to %s Mode", player.Name(), pascalify(ctx.Arguments[0]))
		if player.UUID == ctx.Executor.(*server.PlayerController).UUID {
			msg = fmt.Sprintf("Set own game mode to %s Mode", pascalify(ctx.Arguments[0]))
		}
		ctx.Reply(msg)
	},
}
