package plugins

import (
	"os"

	"github.com/Shopify/go-lua"
	"github.com/dynamitemc/dynamite/logger"
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

func GetLuaVM(logger logger.Logger, plugin *Plugin) *lua.State {
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
	l.SetGlobal("server")

	luaCreateGlobalFunction(l, "Plugin", func(state *lua.State) int {
		if state.IsTable(1) {
			l.Field(1, "identifier")
			identifier, ok := l.ToString(-1)
			if !ok {
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
