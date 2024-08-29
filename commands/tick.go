package commands

import (
	"fmt"
	"time"

	"github.com/zeppelinmc/zeppelin/protocol/text"
	"github.com/zeppelinmc/zeppelin/server"
	"github.com/zeppelinmc/zeppelin/server/command"
)

var tick = command.Command{
	Node: command.NewCommand("tick", command.NewCommand("info"), command.NewCommand("freeze"), command.NewCommand("unfreeze")),
	Callback: func(ccc command.CommandCallContext) {
		tickManager := ccc.Server.(*server.Server).TickManager
		command := ccc.Arguments.Fallback(0, "info")
		num := tickManager.Count()

		switch command {
		case "info":
			freq := tickManager.Frequency()
			ccc.Executor.SystemMessage(text.TextComponent{
				Text: fmt.Sprintf(
					"Server Tickers: %d\nTicking Frequency: %.02ftps (expected ticks per second)",
					num, 1/(float64(freq)/float64(time.Second)),
				),
			})
		case "freeze":
			tickManager.Freeze()
			ccc.Executor.SystemMessage(text.Sprintf("Froze %d tickers", num))
		case "unfreeze":
			tickManager.Unfreeze()
			ccc.Executor.SystemMessage(text.Sprintf("Froze %d tickers", num))
		}

	},
}
