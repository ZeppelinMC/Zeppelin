package core_commands

import (
	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/player"
)

func getServer(executor interface{}) *server.Server {
	if p, ok := executor.(*player.Player); ok {
		return p.Server.(*server.Server)
	} else if c, ok := executor.(*server.Server); ok {
		return c
	}
	return nil
}

var Commands = &commands.Graph{
	Commands: []*commands.Command{
		reload_cmd,
		gamemode_cmd,
		ram_cmd,
		kill_cmd,
		ban_cmd,
		banlist_cmd,
		op_cmd,
		deop_cmd,
		unban_cmd,
		dimension_cmd,
		list_cmd,
		seed_cmd,
		tp_cmd,
		stop_cmd,
		msg_cmd,
		summon_cmd,
		test_cmd,
		nick_cmd,
	},
}
