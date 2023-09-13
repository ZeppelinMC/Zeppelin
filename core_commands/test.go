package core_commands

import (
	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/server/commands"
)

var test_cmd = &commands.Command{
	Name: "test",
	Execute: func(e commands.Executor, s []string) {
		p, _ := e.(*server.PlayerController)
		p.SystemChatMessage("hello!")
	},
}
