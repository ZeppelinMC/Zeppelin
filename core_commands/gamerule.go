package core_commands

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/server/commands"
)

var gamerule_cmd = &commands.Command{
	Name:                "gamerule",
	RequiredPermissions: []string{"server.command.gamerule"},
	Arguments: []commands.Argument{
		commands.NewStrArg("rule", commands.SingleWord).
			SetSuggest(func(ctx commands.SuggestionsContext) {
				srv := ctx.Executor.(*server.PlayerController).Server
				var matches []packet.SuggestionMatch
				for k := range srv.World.Gamerules() {
					matches = append(matches, packet.SuggestionMatch{
						Match: k,
					})
				}
				ctx.Return(matches)
			}),
	},
	Execute: func(ctx commands.CommandContext) {},
}
