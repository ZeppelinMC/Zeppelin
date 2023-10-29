package commands

import (
	"fmt"
	"slices"
	"strings"

	"github.com/aimjel/minecraft/chat"
	"github.com/aimjel/minecraft/packet"
	"github.com/aimjel/minecraft/protocol/types"

	pk "github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/logger"
)

type SuggestionsContext struct {
	Executor      interface{}
	TransactionId int32
	Arguments     []string
	FullCommand   string
}

func (c *SuggestionsContext) Return(suggestions []pk.SuggestionMatch) {
	if p, ok := c.Executor.(interface {
		SendCommandSuggestionsResponse(id int32, start int32, length int32, matches []pk.SuggestionMatch)
	}); ok {
		var start, length int32
		if len(c.Arguments) > 0 {
			arg := c.Arguments[len(c.Arguments)-1]
			start = int32(strings.Index(c.FullCommand, arg))
			length = int32(len(arg))
		} else {
			start = int32(len(c.FullCommand))
			length = int32(len(c.FullCommand))
		}
		p.SendCommandSuggestionsResponse(c.TransactionId, start, length, suggestions)
	}
}

type CommandContext struct {
	Executor           interface{}
	Arguments          []string
	Salt, Timestamp    int64
	ArgumentSignatures []packet.Argument
	FullCommand        string
}

func (ctx *CommandContext) Reply(message chat.Message) {
	if p, ok := ctx.Executor.(interface {
		SystemChatMessage(message chat.Message) error
	}); ok {
		p.SystemChatMessage(message)
	} else {
		fmt.Println(logger.ParseChat(message))
	}
}

func (ctx *CommandContext) Incomplete() {
	ctx.Reply(chat.NewMessage(fmt.Sprintf("§cUnknown or incomplete command, see below for error\n§7%s§r§c§o<--[HERE]", ctx.FullCommand)))
}

func (ctx *CommandContext) ErrorHere(msg string) {
	sp := strings.Split(ctx.FullCommand, " ")
	ctx.Reply(chat.NewMessage(fmt.Sprintf("§c%s\n§7%s §c§n%s§c§o<--[HERE]", msg, strings.Join(sp[:len(sp)-1], " "), sp[len(sp)-1])))
}

func (ctx *CommandContext) Error(msg string) {
	ctx.Reply(chat.NewMessage("§c" + msg))
}

const (
	ChatModeEnabled = iota
	ChatModeCommandsOnly
	ChatModeHidden
)

type Command struct {
	Name                string
	Arguments           []Argument
	Aliases             []string
	Execute             func(ctx CommandContext)
	RequiredPermissions []string
}

type Parser struct {
	ID         int32
	Properties types.CommandProperties
}

type Argument struct {
	Name        string
	Suggest     func(ctx SuggestionsContext)
	Parser      Parser
	Alternative *Argument
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
	packet.Nodes = append(packet.Nodes, types.CommandNode{
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
		packet.Nodes = append(packet.Nodes, types.CommandNode{
			Name:  command.Name,
			Flags: 1 | 0x04,
		})
		for _, argument := range command.Arguments {
			parent := len(packet.Nodes) - 1
			packet.Nodes[parent].Children = append(packet.Nodes[parent].Children, int32(len(packet.Nodes)))
			node := types.CommandNode{Flags: 2, Name: argument.Name, Properties: argument.Parser.Properties, ParserID: argument.Parser.ID}
			if argument.Suggest != nil {
				node.Flags |= 0x10
				node.SuggestionsType = "minecraft:ask_server"
			}
			packet.Nodes = append(packet.Nodes, node)
			if argument.Alternative != nil {
				argument = *argument.Alternative
				packet.Nodes[parent].Children = append(packet.Nodes[parent].Children, int32(len(packet.Nodes)))
				node := types.CommandNode{Flags: 2, Name: argument.Name, Properties: argument.Parser.Properties, ParserID: argument.Parser.ID}
				if argument.Suggest != nil {
					node.Flags |= 0x10
					node.SuggestionsType = "minecraft:ask_server"
				}
				packet.Nodes = append(packet.Nodes, node)
			}
		}
	}
	packet.Nodes[0].Children = rootChildren
	return packet
}

func RegisterCommands(commands ...*Command) *pk.DeclareCommands {
	return Graph{Commands: commands}.Data()
}

func (graph *Graph) FindCommand(name string) (cmd *Command) {
	for _, c := range graph.Commands {
		if c == nil {
			continue
		}
		if c.Name == name {
			cmd = c
			return
		}

		for _, a := range c.Aliases {
			if a == name {
				cmd = c
				return
			}
		}
	}
	return
}

func (graph *Graph) DeleteCommand(name string) (found bool) {
	for i, c := range graph.Commands {
		if c == nil {
			continue
		}
		if c.Name == name {
			graph.Commands = slices.Delete(graph.Commands, i, i+1)
			return true
		}

		for _, a := range c.Aliases {
			if a == name {
				graph.Commands = slices.Delete(graph.Commands, i, i+1)
				return true
			}
		}
	}
	return false
}
