// Package command provides utilities for handling and registering commands
package command

import (
	"github.com/zeppelinmc/zeppelin/protocol/text"
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
	Server   any
	Executor Caller

	Arguments Arguments
}

func (c CommandCallContext) Reply(msg text.TextComponent) error {
	return c.Executor.SystemMessage(msg)
}
