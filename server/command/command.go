// Package command provides utilities for handling and registering commands
package command

import (
	"github.com/zeppelinmc/zeppelin/server/session"
)

type Command struct {
	Node      Node
	Aliases   []string
	Namespace string

	Callback           func(CommandCallContext)
	SuggestionCallback func()
}

type Arguments []string

func (a Arguments) At(i int) string {
	if i < 0 || len(a) <= i {
		return ""
	}
	return a[i]
}

func (a Arguments) Fallback(i int, fb string) string {
	if i < 0 || len(a) <= i {
		return fb
	}
	return a[i]
}

type CommandCallContext struct {
	Command  Command
	Executor session.Session
	Server   any

	Arguments Arguments
}
