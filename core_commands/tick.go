package core_commands

import (
	"fmt"
	"time"

	"github.com/zeppelinmc/zeppelin/server"
	"github.com/zeppelinmc/zeppelin/server/command"
	"github.com/zeppelinmc/zeppelin/text"
)

var tick = command.Command{
	Node: command.NewCommand("tick", command.NewCommand("info")),
	Callback: func(ccc command.CommandCallContext) {
		num := ccc.Server.(*server.Server).TickManager.Count()
		frequency := ccc.Server.(*server.Server).TickManager.Frequency()

		ccc.Executor.SystemMessage(text.TextComponent{
			Text: fmt.Sprintf(
				"Server Tickers: %d\nTicking Frequency: %.02ftps (expected ticks per second)",
				num, 1/(float64(frequency)/float64(time.Second)),
			),
		})
	},
}
