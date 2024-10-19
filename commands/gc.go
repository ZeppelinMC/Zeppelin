package commands

import (
	"runtime"

	"github.com/zeppelinmc/zeppelin/protocol/text"
	"github.com/zeppelinmc/zeppelin/server/command"
)

var gc = command.Command{
	Node:      command.NewLiteral("gc"),
	Namespace: "zeppelin",
	Callback: func(ccc command.CommandCallContext) {
		runtime.GC()

		ccc.Executor.SystemMessage(text.Text("Done.").WithColor(text.BrightGreen))
	},
}
