package core_commands

import (
	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/server/commands"
)

func getServer(executor interface{}) *server.Server {
	if p, ok := executor.(*server.Session); ok {
		return p.Server
	} else if c, ok := executor.(*server.ConsoleExecutor); ok {
		return c.Server
	}
	return nil
}

var Commands = &commands.Graph{
	Commands: []*commands.Command{
		test_cmd,
		reload_cmd,
		gamemode_cmd,
		ram_cmd,
		kill_cmd,
		gamerule_cmd,
		ban_cmd,
		banlist_cmd,
		op_cmd,
		deop_cmd,
		unban_cmd,
		dimension_cmd,
		list_cmd,
		seed_cmd,
		tp_cmd,
	},
}
