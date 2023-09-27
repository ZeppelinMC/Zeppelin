package commands

import (
	"fmt"
	"strings"

	pk "github.com/aimjel/minecraft/packet"
)

type CommandContext struct {
	Executor interface {
		SystemChatMessage(s string) error
	} `js:"executor"`
	Arguments   []string `js:"arguments"`
	FullCommand string   `js:"fullCommand"`
}

func (ctx *CommandContext) Reply(content string) {
	ctx.Executor.SystemChatMessage(content)
}

func (ctx *CommandContext) Incomplete() {
	ctx.Reply(fmt.Sprintf("§cUnknown or incomplete command, see below for error\n§7%s§r§c§o<--[HERE]", ctx.FullCommand))
}

func (ctx *CommandContext) ErrorAt(msg string) {
	sp := strings.Split(ctx.FullCommand, " ")
	ctx.Reply(fmt.Sprintf("§c%s\n§7%s §c§n%s§c§o<--[HERE]", msg, strings.Join(sp[:len(sp)-1], " "), sp[len(sp)-1]))
}

func (ctx *CommandContext) Error(msg string) {
	ctx.Reply("§c" + msg)
}

const (
	ChatModeEnabled = iota
	ChatModeCommandsOnly
	ChatModeHidden
)

type Command struct {
	Name                string                   `js:"name"`
	Arguments           []Argument               `js:"arguments"`
	Aliases             []string                 `js:"aliases"`
	Execute             func(ctx CommandContext) `js:"execute"`
	RequiredPermissions []string                 `js:"requiredPermissions"`
}

type Properties struct {
	Flags      uint8
	Min, Max   uint64
	Identifier string
}

type Parser struct {
	ID         int32      `js:"id"`
	Properties Properties `js:"properties"`
}

type Argument struct {
	Name           string `js:"name"`
	SuggestionType string `js:"suggestionType"`
	Parser         Parser `js:"parser"`
}

type Graph struct {
	Commands []*Command
}

func (graph *Graph) AddCommands(commands ...*Command) *Graph {
	graph.Commands = append(graph.Commands, commands...)
	return graph
}

func (command *Command) AddArguments(arguments ...Argument) *Command {
	command.Arguments = append(command.Arguments, arguments...)
	return command
}

func (graph Graph) Data() *pk.DeclareCommands {
	packet := &pk.DeclareCommands{}
	packet.Nodes = append(packet.Nodes, pk.Node{
		Flags: 0,
	})
	commands := graph.Commands
	rootChildren := []int32{}
	for _, command := range commands {
		if command == nil {
			continue
		}
		for _, alias := range command.Aliases {
			commands = append(commands, &Command{
				Name:      alias,
				Arguments: command.Arguments,
			})
		}
	}
	for _, command := range commands {
		if command == nil {
			continue
		}
		rootChildren = append(rootChildren, int32(len(packet.Nodes)))
		packet.Nodes = append(packet.Nodes, pk.Node{
			Name:  command.Name,
			Flags: 1,
		})
		for _, argument := range command.Arguments {
			parent := len(packet.Nodes) - 1
			packet.Nodes[parent].Children = append(packet.Nodes[parent].Children, int32(len(packet.Nodes)))
			node := pk.Node{Flags: 2, Name: argument.Name, Properties: argument.Parser.Properties, ParserID: argument.Parser.ID}
			if argument.SuggestionType != "" {
				node.Flags |= 0x10
				node.SuggestionsType = argument.SuggestionType
			}
			packet.Nodes = append(packet.Nodes, node)
		}
	}
	packet.Nodes[0].Children = rootChildren
	return packet
}

func RegisterCommands(commands ...*Command) *pk.DeclareCommands {
	return Graph{Commands: commands}.Data()
}
