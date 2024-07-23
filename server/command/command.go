package command

import (
	"github.com/zeppelinmc/zeppelin/server/session"
)

type Command struct {
	Name      string
	Aliases   []string
	Arguments []Argument

	Callback           func(CommandCallContext)
	SuggestionCallback func()
}

type CommandCallContext struct {
	Command  Command
	Executor session.Session
	Server   any

	Arguments []any
}

type Argument struct {
	Name string
}
