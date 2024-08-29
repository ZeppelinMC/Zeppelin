package commands

import (
	"fmt"
	"runtime"

	"github.com/zeppelinmc/zeppelin/protocol/net/io/buffers"
	"github.com/zeppelinmc/zeppelin/protocol/text"
	"github.com/zeppelinmc/zeppelin/server"
	"github.com/zeppelinmc/zeppelin/server/command"
)

var memStats runtime.MemStats

var mem = command.Command{
	Node:      command.NewCommand("mem"),
	Namespace: "zeppelin",
	Callback: func(ccc command.CommandCallContext) {
		runtime.ReadMemStats(&memStats)
		loaded := ccc.Server.(*server.Server).World.LoadedChunks()
		goroutines := runtime.NumGoroutine()

		ccc.Executor.SystemMessage(text.TextComponent{
			Text: fmt.Sprintf(
				"Server memory usage: \n\nAlloc: %dMiB, Total Alloc: %dMiB\nLoaded Chunks: %d\nBuffer size: %dB\nGoroutines: %d",
				memStats.Alloc/1024/1024, memStats.TotalAlloc/1024/1024, loaded, buffers.Size(), goroutines,
			),
		})
	},
}
