package commands

import (
	pk "github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/player"
)

type Command struct {
	Name      string
	Arguments []Argument
	Aliases   []string
	Execute   func(*player.Player, []string)
}

type Properties struct {
	Flags      uint8  `json:",omitempty"`
	Min, Max   uint64 `json:",omitempty"`
	Identifier string `json:",omitempty"`
}

type Argument struct {
	Name           string
	ParserID       int32
	SuggestionType string
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
