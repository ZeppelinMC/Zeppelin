package core_commands

import (
	"fmt"
	"strings"

	"github.com/aimjel/minecraft/chat"
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
		commands.NewGamemodeArgument("mode"),
		commands.NewEntityArgument("player", commands.EntityPlayerOnly),
	},
	Execute: func(ctx commands.CommandContext) {
		if len(ctx.Arguments) == 0 {
			ctx.Incomplete()
			return
		}
		gm := p.Gamemode(ctx.Arguments[0])
		if gm == -1 {
			ctx.ErrorHere(fmt.Sprintf("Unknown game mode: %s", ctx.Arguments[0]))
			return
		}
		var player *server.Session
		if len(ctx.Arguments) == 1 {
			if p, ok := ctx.Executor.(*server.Session); !ok {
				ctx.Incomplete()
				return
			} else {
				player = p
			}
		} else {
			p := getServer(ctx.Executor).FindPlayer(ctx.Arguments[1])
			if p == nil {
				ctx.Error("No player was found")
				return
			}
			player = p
		}
		if int(player.Player.GameMode()) == gm {
			return
		}
		player.SetGameMode(byte(gm))
		msg := player.Server.Translate("commands.gamemode.success.other", chat.NewMessage(player.Name()), chat.NewMessage(pascalify(ctx.Arguments[0])))
		if exe, ok := ctx.Executor.(*server.Session); ok && player.UUID == exe.UUID {
			msg = player.Server.Translate("commands.gamemode.success.self", chat.NewMessage(pascalify(ctx.Arguments[0])))
		}
		ctx.Reply(msg)
	},
}
