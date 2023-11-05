package commands

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/aimjel/minecraft/chat"
	pk "github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/logger"
)

type CommandContext struct {
	Command            *Command
	Executor           interface{}
	Arguments          []string
	Salt, Timestamp    int64
	ArgumentSignatures []pk.Argument
	FullCommand        string
}

func findArgument(a []Argument, n string) int {
	for i, arg := range a {
		if arg.Name == n {
			return i
		}
	}
	return -1
}

func (ctx CommandContext) GetString(name string) (value string, ok bool) {
	for i, a := range ctx.Arguments {
		arg := ctx.Command.Arguments[i]
		if (arg.Parser.ID >= 8 && arg.Parser.ID <= 11) || arg.Parser.ID == 27 {
			ctx.Arguments = slices.Delete(ctx.Arguments, i+1, i+3)
		}
		if arg.Name == name {
			return a, true
		}
	}
	return "", false
}

func (ctx CommandContext) GetInt32(name string) (value int32, ok bool) {
	s, ok := ctx.GetString(name)
	if !ok {
		return 0, false
	}
	i, e := strconv.ParseInt(s, 10, 32)
	return int32(i), e == nil
}

func (ctx CommandContext) GetInt64(name string) (value int64, ok bool) {
	s, ok := ctx.GetString(name)
	if !ok {
		return 0, false
	}
	i, e := strconv.ParseInt(s, 10, 64)
	return i, e == nil
}

func (ctx CommandContext) GetFloat32(name string) (value float32, ok bool) {
	s, ok := ctx.GetString(name)
	if !ok {
		return 0, false
	}
	i, e := strconv.ParseFloat(s, 32)
	return float32(i), e == nil
}

func (ctx CommandContext) GetFloat64(name string) (value float64, ok bool) {
	s, ok := ctx.GetString(name)
	if !ok {
		return 0, false
	}
	i, e := strconv.ParseFloat(s, 64)
	return i, e == nil
}

func (ctx CommandContext) GetBool(name string) (value bool, ok bool) {
	s, ok := ctx.GetString(name)
	if !ok {
		return false, false
	}
	i, e := strconv.ParseBool(s)
	return i, e == nil
}

func (ctx *CommandContext) Reply(message chat.Message) {
	if p, ok := ctx.Executor.(interface {
		SystemChatMessage(message chat.Message) error
	}); ok {
		p.SystemChatMessage(message)
	} else {
		fmt.Print(strings.ReplaceAll(logger.ParseChat(message), "\n", "\n\r"))
		fmt.Print("\n\r> ")
	}
}

func (ctx *CommandContext) Incomplete() {
	_, ok := ctx.Executor.(interface {
		SystemChatMessage(message chat.Message) error
	})
	ctx.Reply(chat.NewMessage(fmt.Sprintf("§cUnknown or incomplete command, see below for error"+cond(ok, "", "\r")+"\n§7%s§r§c§o<--[HERE]", ctx.FullCommand)))
}

func (ctx *CommandContext) ErrorHere(msg string) {
	_, ok := ctx.Executor.(interface {
		SystemChatMessage(message chat.Message) error
	})
	sp := strings.Split(ctx.FullCommand, " ")
	ctx.Reply(chat.NewMessage(fmt.Sprintf("§c%s\n"+cond(ok, "", "\r")+"§7%s §c§n%s§c§o<--[HERE]", msg, strings.Join(sp[:len(sp)-1], " "), sp[len(sp)-1])))
}

func (ctx *CommandContext) Error(msg string) {
	ctx.Reply(chat.NewMessage("§c" + msg))
}
