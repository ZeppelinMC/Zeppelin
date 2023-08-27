package commands

import pk "github.com/aimjel/minecraft/packet"

const (
	StringSingleWord = iota
	StringQuotablePhrase
	StringGreedyPhrase
)

type Command struct {
	Name      string
	Arguments []Argument
	Aliases   []string
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

func NewGraph(commands ...*Command) *Graph {
	graph := &Graph{Commands: commands}
	return graph
}

func (graph *Graph) AddCommands(commands ...*Command) *Graph {
	graph.Commands = append(graph.Commands, commands...)
	return graph
}

func (command *Command) AddArguments(arguments ...Argument) *Command {
	command.Arguments = append(command.Arguments, arguments...)
	return command
}

func NewCommand(name string, arguments ...Argument) *Command {
	return &Command{Name: name, Arguments: arguments}
}

func NewBoolArgument(name string) Argument {
	return Argument{
		Name:     name,
		ParserID: 0,
	}
}

func NewFloatArgument(name string, properties struct {
	Min *uint64
	Max *uint64
}) Argument {
	props := Properties{Flags: 0}
	if properties.Min != nil {
		props.Flags |= 1
		props.Min = *properties.Min
	}
	if properties.Max != nil {
		props.Flags |= 2
		props.Max = *properties.Max
	}
	return Argument{
		Name:       name,
		ParserID:   1,
		Properties: props,
	}
}

func NewDoubleArgument(name string, properties struct {
	Min *uint64
	Max *uint64
}) Argument {
	props := Properties{Flags: 0}
	if properties.Min != nil {
		props.Flags |= 1
		props.Min = *properties.Min
	}
	if properties.Max != nil {
		props.Flags |= 2
		props.Max = *properties.Max
	}
	return Argument{
		Name:       name,
		ParserID:   2,
		Properties: props,
	}
}

func NewIntegerArgument(name string, properties struct {
	Min *int64
	Max *int64
}) Argument {
	props := Properties{Flags: 0}
	if properties.Min != nil {
		props.Flags |= 1
		props.Min = uint64(*properties.Min)
	}
	if properties.Max != nil {
		props.Flags |= 2
		props.Max = uint64(*properties.Max)
	}
	return Argument{
		Name:       name,
		ParserID:   3,
		Properties: props,
	}
}

func NewLongArgument(name string, properties struct {
	Min *int64
	Max *int64
}) Argument {
	props := Properties{Flags: 0}
	if properties.Min != nil {
		props.Flags |= 1
		props.Min = uint64(*properties.Min)
	}
	if properties.Max != nil {
		props.Flags |= 2
		props.Max = uint64(*properties.Max)
	}
	return Argument{
		Name:       name,
		ParserID:   4,
		Properties: props,
	}
}

func NewStringArgument(name string, properties byte) Argument {
	props := Properties{Flags: properties}
	return Argument{
		Name:       name,
		ParserID:   5,
		Properties: props,
	}
}

func (graph Graph) Data(packet pk.DeclareCommands) {
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
}
