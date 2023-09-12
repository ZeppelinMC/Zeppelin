package core_commands

import (
	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/player"
)

var test_cmd = &commands.Command{
	Name: "test",
	Execute: func(p *player.Player, s []string) {
		p.SystemChatMessage("hello!")
	},
}
