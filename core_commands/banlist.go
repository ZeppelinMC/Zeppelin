package core_commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dynamitemc/dynamite/server/commands"
)

func point[T any](t T) *T {
	return &t
}

var banlist_cmd = &commands.Command{
	Name:                "banlist",
	RequiredPermissions: []string{"server.command.banlist"},
	Arguments: []commands.Argument{
		commands.NewIntegerArgument("page", struct {
			Min *int64
			Max *int64
		}{
			Min: point[int64](1),
		}),
	},
	Execute: func(ctx commands.CommandContext) {
		server := getServer(ctx.Executor)
		page := 1
		if len(ctx.Arguments) > 0 {
			if i, err := strconv.Atoi(ctx.Arguments[0]); err == nil {
				page = i
			}
		}
		_ = fmt.Sprint(page) // todo add paging
		str := "§lBan list:\n"
		for i, b := range server.BannedPlayers {
			str += fmt.Sprintf("§l%d§r%s", i+1, b.Name)
			if b.Created != "" {
				d, _ := time.Parse("2006-01-02T15:04:05Z07:00", b.Created)
				str += " - " + d.Format("Mon Jan _2 15:04:05 2006")
			}
		}
		ctx.Reply(str)
	},
}
