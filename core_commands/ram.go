package core_commands

import (
	"fmt"
	"runtime"

	"github.com/dynamitemc/dynamite/server/commands"
)

var ram_cmd = &commands.Command{
	Name:                "ram",
	Aliases:             []string{"mem"},
	RequiredPermissions: []string{},
	Execute: func(ctx commands.CommandContext) {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		srv := getServer(ctx.Executor)
		ctx.Reply(fmt.Sprintf("%d MiB memory used!\nLoaded chunks: %d", bytesToMegabytes(m.Alloc), srv.World.Overworld().LoadedChunks()))
	},
}

func bytesToMegabytes(b uint64) uint64 {
	return b / 1024 / 1024
}
