package core_commands

import (
	"fmt"
	"runtime"

	"github.com/zeppelinmc/zeppelin/server/command"
	"github.com/zeppelinmc/zeppelin/text"
)

var memStats runtime.MemStats

var MemCommand = command.Command{
	Name: "mem",
	Callback: func(ccc command.CommandCallContext) {
		runtime.ReadMemStats(&memStats)

		ccc.Executor.SystemMessage(text.TextComponent{
			Text: fmt.Sprintf("Server memory usage: \n\nAlloc: %dMiB, Total Alloc: %dMiB", memStats.Alloc/1024/1024, memStats.TotalAlloc/1024/1024),
		})
	},
}
