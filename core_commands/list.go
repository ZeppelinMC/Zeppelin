package core_commands

import (
	"fmt"

	"github.com/aimjel/minecraft/chat"
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/commands"
)

var list_cmd = &commands.Command{
	Name:                "list",
	RequiredPermissions: []string{"server.command.list"},
	Arguments: []commands.Argument{
		commands.NewStrArg("uuids", commands.SingleWord).SetSuggest(func(ctx commands.SuggestionsContext) {
			ctx.Return([]packet.SuggestionMatch{
				{Match: "uuids"},
			})
		}),
	},
	Execute: func(ctx commands.CommandContext) {
		srv := getServer(ctx.Executor)
		players := srv.Players()
		if len(players) == 0 {
			ctx.Reply(chat.NewMessage("No players online"))
			return
		}
		msg := fmt.Sprintf("There are %d of a max of %d players online: ", len(players), srv.Config.MaxPlayers)
		var index int
		for _, p := range players {
			msg += p.Name()
			if len(ctx.Arguments) == 1 && ctx.Arguments[0] == "uuids" {
				msg += fmt.Sprintf(" (%s)", p.UUID())
			}
			if index != len(players)-1 {
				msg += ", "
			}
			index++
		}
		ctx.Reply(chat.NewMessage(msg))
	},
}
