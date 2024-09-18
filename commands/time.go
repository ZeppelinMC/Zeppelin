package commands

import (
	"strconv"

	"github.com/zeppelinmc/zeppelin/protocol/text"
	"github.com/zeppelinmc/zeppelin/server"
	"github.com/zeppelinmc/zeppelin/server/command"
)

var timecmd = command.Command{
	Node: command.NewLiteral("time" /*command.NewCommand("add", command.NewTimeArgument("time", 0)), command.NewCommand("set", command.NewTimeArgument("time", 0))*/),
	Callback: func(ccc command.CommandCallContext) {
		command := ccc.Arguments.At(0)
		w := ccc.Server.(*server.Server).World

		switch command {
		case "set":
			t := ccc.Arguments.At(1)
			time, err := strconv.Atoi(t)
			if t == "" || err != nil {
				ccc.Reply(text.Sprint("Invalid time"))
				return
			}

			a, _ := w.Time()
			ccc.Executor.UpdateTime(a, int64(time))
		}
	},
}
