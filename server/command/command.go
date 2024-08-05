package command

import (
	"github.com/zeppelinmc/zeppelin/server/session"
)

type Command struct {
	Node    Node
	Aliases []string

	Callback           func(CommandCallContext)
	SuggestionCallback func()
}

type CommandCallContext struct {
	Command  Command
	Executor session.Session
	Server   any

	Arguments []any
}
