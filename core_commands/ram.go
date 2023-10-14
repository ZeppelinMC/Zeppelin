package core_commands

import (
	"fmt"
	"runtime"

	"github.com/dynamitemc/dynamite/server/commands"
)

var ram_cmd = &commands.Command{
	Name:    "ram",
	Aliases: []string{"mem"},
	RequiredPermissions: []string{
		"server.command.ram",
	},
	Execute: func(ctx commands.CommandContext) {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		ctx.Reply(fmt.Sprintf("Allocated: %v MiB, Total Allocated: %v MiB, Heap in Use: %v MiB", bytesToMegabytes(m.Alloc), bytesToMegabytes(m.TotalAlloc), bytesToMegabytes(m.HeapInuse)))
	},
}

func bytesToMegabytes(b uint64) uint64 {
	return b / 1024 / 1024
}
