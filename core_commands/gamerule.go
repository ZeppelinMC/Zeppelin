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
		{
			Name: "rule",
			Parser: commands.Parser{
				ID: 5,
				Properties: commands.Properties{
					Flags: commands.StringSingleWord,
				},
			},
			Suggest: func(ctx commands.SuggestionsContext) {
				srv := ctx.Executor.(*server.PlayerController).Server
				var matches []packet.SuggestionMatch
				for k := range srv.World.Gamerules() {
					matches = append(matches, packet.SuggestionMatch{
						Match: k,
					})
				}
				ctx.Return(matches)
			},
		},
	},
	Execute: func(ctx commands.CommandContext) {},
}
