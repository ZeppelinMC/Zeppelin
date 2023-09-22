package server

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	lua "github.com/Shopify/go-lua"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/robertkrimen/otto"
)

var unrecognizedScript = errors.New("unrecognized scripting language")

const (
	PluginTypeJavaScript = iota
	PluginTypeLua
)

type Plugin struct {
	Identifier string
	Filename   string
	JSLoader   *otto.Otto
	LuaLoader  *lua.State
	Type       int
}

func luaCreateFunction(l *lua.State, k string, f lua.Function) {
	l.PushGoFunction(f)
	l.SetField(-2, k)
}

func luaCreateGlobalFunction(l *lua.State, k string, f lua.Function) {
	l.PushGoFunction(f)
	l.SetGlobal(k)
}

func getLuaVM(logger logger.Logger, plugin *Plugin) *lua.State {
	l := lua.NewState()
	l.NewTable()
	luaCreateFunction(l, "close", func(state *lua.State) int {
		code := 0
		if c, ok := state.ToInteger(1); ok {
			code = c
		}
		os.Exit(code)
		return 0
	})
	l.NewTable()
	luaCreateFunction(l, "info", func(state *lua.State) int {
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
	})
	luaCreateFunction(l, "error", func(state *lua.State) int {
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
	})
	luaCreateFunction(l, "debug", func(state *lua.State) int {
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
	})
	luaCreateFunction(l, "warn", func(state *lua.State) int {
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
		} else {
			logger.Error("Failed to load plugin %s: invalid plugin data", plugin.Filename)
		}
		return 0
	})
	return l
}

func getJavaScriptVM(logger logger.Logger, plugin *Plugin) *otto.Otto {
	vm := otto.New()
	server, _ := vm.Object("server = {}")
	log, _ := vm.Object("server.logger = {}")

	server.Set("close", func(call otto.FunctionCall) otto.Value {
		var code int64 = 0
		if c := call.Argument(0); c.IsNumber() {
			code, _ = c.ToInteger()
		}
		os.Exit(int(code))
		return otto.UndefinedValue()
	})

	log.Set("info", func(call otto.FunctionCall) otto.Value {
		if s, err := call.Argument(0).ToString(); err != nil {
			return otto.UndefinedValue()
		} else {
			var data []interface{}
			for _, a := range call.ArgumentList[1:] {
				data = append(data, a)
			}
			logger.Info(s, data...)
			return otto.UndefinedValue()
		}
	})

	log.Set("debug", func(call otto.FunctionCall) otto.Value {
		if s, err := call.Argument(0).ToString(); err != nil {
			return otto.UndefinedValue()
		} else {
			var data []interface{}
			for _, a := range call.ArgumentList[1:] {
				data = append(data, a)
			}
			logger.Debug(s, data...)
			return otto.UndefinedValue()
		}
	})

	log.Set("warn", func(call otto.FunctionCall) otto.Value {
		if s, err := call.Argument(0).ToString(); err != nil {
			return otto.UndefinedValue()
		} else {
			var data []interface{}
			for _, a := range call.ArgumentList[1:] {
				data = append(data, a)
			}
			logger.Warn(s, data...)
			return otto.UndefinedValue()
		}
	})

	log.Set("error", func(call otto.FunctionCall) otto.Value {
		if s, err := call.Argument(0).ToString(); err != nil {
			return otto.UndefinedValue()
		} else {
			var data []interface{}
			for _, a := range call.ArgumentList[1:] {
				data = append(data, a)
			}
			logger.Error(s, data...)
			return otto.UndefinedValue()
		}
	})

	vm.Set("Plugin", func(call otto.FunctionCall) otto.Value {
		if obj := call.Argument(0); obj.IsObject() {
			data := obj.Object()
			identifier, _ := data.Get("identifier")
			if identifier.IsUndefined() {
				logger.Error("Failed to load plugin %s: identifier was not specified", plugin.Filename)
			}
			plugin.Identifier, _ = identifier.ToString()
		} else {
			logger.Error("Failed to load plugin %s: invalid plugin data", plugin.Filename)
		}
		return otto.UndefinedValue()
	})
	return vm
}

func (srv *Server) LoadPlugins() error {
	err := os.Mkdir("plugins", 0755)
	if err != nil {
		if !errors.Is(err, fs.ErrExist) {
			return err
		}
	}
	dir, err := os.ReadDir("plugins")
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, file := range dir {
		if plugin, err := srv.LoadPlugin("plugins/" + file.Name()); err != nil {
			return err
		} else {
			if plugin == nil {
				continue
			}
			srv.Plugins[plugin.Identifier] = plugin
			srv.Logger.Info("Finished loading plugin %s", plugin.Identifier)
		}
	}
	return nil
}

func (srv *Server) LoadPlugin(path string) (*Plugin, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	sp := strings.Split(path, "/")
	filename := sp[len(sp)-1]
	switch {
	case strings.HasSuffix(path, ".js"):
		{
			plugin := &Plugin{Filename: filename, Type: PluginTypeJavaScript}
			js := getJavaScriptVM(srv.Logger, plugin)
			plugin.JSLoader = js
			_, err := js.Run(string(file))
			if err != nil {
				srv.Logger.Error("Failed to load plugin %s: %s", filename, err)
			}
			return plugin, nil
		}
	case strings.HasSuffix(path, ".lua"):
		{
			plugin := &Plugin{Filename: filename, Type: PluginTypeJavaScript}
			l := getLuaVM(srv.Logger, plugin)
			plugin.LuaLoader = l
			lua.DoString(l, string(file))
			return plugin, nil
		}
	}
	return nil, nil
}
