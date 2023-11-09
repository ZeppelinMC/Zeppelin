package handler

import (
	"strings"

	"github.com/dynamitemc/dynamite/server/commands"
	"github.com/dynamitemc/dynamite/server/player"
)

func CommandSuggestionsRequest(id int32, content string, graph *commands.Graph, state *player.Player) {
	args := strings.Split(strings.TrimSpace(content), " ")
	cmd := strings.TrimPrefix(args[0], "/")
	var command *commands.Command
	var argument *commands.Argument
	for _, c := range graph.Commands {
		if c == nil {
			continue
		}
		if c.Name == cmd {
			command = c
		}

		for _, a := range c.Aliases {
			if a == cmd {
				command = c
			}
		}
	}
	if command == nil || !state.HasPermissions(command.RequiredPermissions) {
		return
	}
	index := len(args[1:])
	if len(command.Arguments) <= index {
		index = len(command.Arguments)
	}
	if len(command.Arguments) <= index {
		return
	}
	argument = &command.Arguments[index]
	if argument == nil {
		return
	}
	ctx := commands.SuggestionsContext{
		Arguments:     args[1:],
		Executor:      state,
		FullCommand:   content,
		TransactionId: id,
	}
	argument.Suggest(ctx)
}
