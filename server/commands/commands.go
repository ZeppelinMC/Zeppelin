package commands

import (
	pk "github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/player"
)

type CommandContext struct {
	Executor interface {
		SystemChatMessage(s string) error
		OnGround() bool
		ClientSettings() player.ClientInformation
		Position() (x float64, y float64, z float64)
		Rotation() (yaw float32, pitch float32)
	} `js:"executor"`
	Arguments []string `js:"arguments"`
}

func (ctx *CommandContext) Reply(content string) {
	ctx.Executor.SystemChatMessage(content)
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
	Flags      uint8  `json:",omitempty"`
	Min, Max   uint64 `json:",omitempty"`
	Identifier string `json:",omitempty"`
}

type Argument struct {
	Name           string `js:"name"`
	ParserID       int32  `js:"parserId"`
	SuggestionType string `js:"suggestionType"`
	Properties     Properties
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
		for _, alias := range command.Aliases {
			commands = append(commands, &Command{
				Name:      alias,
				Arguments: command.Arguments,
			})
		}
	}
	for _, command := range commands {
		rootChildren = append(rootChildren, int32(len(packet.Nodes)))
		parent := len(packet.Nodes)
		packet.Nodes = append(packet.Nodes, pk.Node{
			Name:  command.Name,
			Flags: 1,
		})
		for _, argument := range command.Arguments {
			packet.Nodes[parent].Children = append(packet.Nodes[parent].Children, int32(len(packet.Nodes)))
			node := pk.Node{Flags: 2, Name: argument.Name, Properties: argument.Properties, ParserID: argument.ParserID}
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
