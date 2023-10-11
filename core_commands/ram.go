package core_commands

import (
	"fmt"
	"runtime"

	"github.com/dynamitemc/dynamite/server/commands"
)

var ram_cmd = &commands.Command{
	Name:    "ram",
	Aliases: []string{"mem"},
	Execute: func(ctx commands.CommandContext) {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		ctx.Reply(fmt.Sprintf("Allocated: %v MiB, Total Allocated: %v MiB, Heap in Use: %v MiB", bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.HeapInuse)))
	},
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
