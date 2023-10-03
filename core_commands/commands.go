package core_commands

import "github.com/dynamitemc/dynamite/server/commands"

var Commands = &commands.Graph{
	Commands: []*commands.Command{
		test_cmd,
		reload_cmd,
		gamemode_cmd,
		restart_cmd,
	},
}
