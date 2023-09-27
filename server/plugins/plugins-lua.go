package plugins

import (
	"fmt"
	"os"
	"slices"

	"github.com/Shopify/go-lua"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server/commands"
)

func luaCreateFunction(l *lua.State, k string, f lua.Function) {
	l.PushGoFunction(f)
	l.SetField(-2, k)
}

func luaCreateGlobalFunction(l *lua.State, k string, f lua.Function) {
	l.PushGoFunction(f)
	l.SetGlobal(k)
}

func luaTable(l *lua.State, s [][2]interface{}) {
	l.NewTable()
	for _, t := range s {
		key, ok := t[0].(string)
		if !ok {
			continue
		}
		switch f := t[1].(type) {
		case func(state *lua.State) int:
			{
				luaCreateFunction(l, key, f)
			}
		}
	}
}

func luaGlobalTable(l *lua.State, s [][2]interface{}) {
	l.NewTable()
	for _, t := range s {
		key, ok := t[0].(string)
		if !ok {
			continue
		}
		switch f := t[1].(type) {
		case func(state *lua.State) int:
			{
				luaCreateFunction(l, key, f)
			}
		}
	}
}

func GetLuaVM(srv interface {
	GetCommandGraph() *commands.Graph
}, logger logger.Logger, plugin *Plugin) *lua.State {
	graph := srv.GetCommandGraph()
	l := lua.NewState()
	luaGlobalTable(l, [][2]interface{}{
		{
			"close",
			func(state *lua.State) int {
				code := 0
				if c, ok := state.ToInteger(1); ok {
					code = c
				}
				os.Exit(code)
				return 0
			},
		},
	})
	luaTable(l, [][2]interface{}{
		{
			"info",
			func(state *lua.State) int {
				text, ok := state.ToString(1)
				if !ok {
					return 0
				}
				var data []interface{}
				for i := 2; ; i++ {
					val := state.ToValue(i)
					if val == nil {
						break
					}
					data = append(data, val)
				}
				logger.Info(text, data...)
				return 0
			},
		},
		{
			"error",
			func(state *lua.State) int {
				text, ok := state.ToString(1)
				if !ok {
					return 0
				}
				var data []interface{}
				for i := 2; ; i++ {
					val := state.ToValue(i)
					if val == nil {
						break
					}
					data = append(data, val)
				}
				logger.Error(text, data...)
				return 0
			},
		},
		{
			"debug",
			func(state *lua.State) int {
				text, ok := state.ToString(1)
				if !ok {
					return 0
				}
				var data []interface{}
				for i := 2; ; i++ {
					val := state.ToValue(i)
					if val == nil {
						break
					}
					data = append(data, val)
				}
				logger.Debug(text, data...)
				return 0
			},
		},
		{
			"warn",
			func(state *lua.State) int {
				text, ok := state.ToString(1)
				if !ok {
					return 0
				}
				var data []interface{}
				for i := 2; ; i++ {
					val := state.ToValue(i)
					if val == nil {
						break
					}
					data = append(data, val)
				}
				logger.Warn(text, data...)
				return 0
			},
		},
	})
	l.SetField(-2, "logger")
	luaTable(l, [][2]interface{}{
		{
			"register",
			func(state *lua.State) int {
				if !state.IsTable(1) {
					logger.Error("Plugin error: %s::server.commands.set argument at 0 is not an object", plugin.Identifier)
					return 0
				}
				l.Field(1, "name")
				name, ok := l.ToString(-1)
				if !ok || name == "" {
					logger.Error("Plugin error: %s::server.commands.set argument at 0: command name was not specified", plugin.Identifier)
				}
				graph.Commands = append(graph.Commands, &commands.Command{
					Name: name,
				})
				return 0
			},
		},
		{
			"delete",
			func(state *lua.State) int {
				name, ok := state.ToString(1)
				if !ok {
					logger.Error("Plugin error: %s::server.commands.delete argument at 0 is not a string", plugin.Identifier)
					return 0
				}
				for i, cmd := range graph.Commands {
					if cmd.Name == name {
						graph.Commands = slices.Delete(graph.Commands, i, i+1)
						return 0
					}
				}
				return 0
			},
		},
	})
	l.Field(-2, "commands")
	l.SetGlobal("server")

	luaCreateGlobalFunction(l, "Plugin", func(state *lua.State) int {
		fmt.Println("h")
		if state.IsTable(1) {
			l.Field(1, "identifier")
			identifier, ok := l.ToString(-1)
			if !ok || identifier == "" {
				logger.Error("Failed to load plugin %s: identifier was not specified", plugin.Filename)
			}
			plugin.Identifier = identifier
			plugin.Initialized = true
		} else {
			logger.Error("Failed to load plugin %s: invalid plugin data", plugin.Filename)
		}
		return 0
	})
	return l
}
