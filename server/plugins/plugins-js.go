package plugins

import (
	"os"
	"slices"

	"github.com/dop251/goja"
	"github.com/dynamitemc/dynamite/logger"
	"github.com/dynamitemc/dynamite/server/commands"
)

type pluginConfiguration struct {
	Identifier string `js:"identifier"`
	OnLoad     func() `js:"onLoad"`
}

func GetJavaScriptVM(srv interface {
	GetCommandGraph() *commands.Graph
}, logger logger.Logger, plugin *Plugin, root string) *goja.Runtime {
	vm := goja.New()
	vm.SetFieldNameMapper(goja.TagFieldNameMapper("js", true))
	server := vm.NewObject()
	log := vm.NewObject()

	server.Set("close", func(c int64) {
		var code int64 = c
		os.Exit(int(code))
	})

	graph := srv.GetCommandGraph()
	cmds := vm.NewObject()
	cmds.Set("register", func(cmds ...*commands.Command) {
		graph.AddCommands(cmds...)
	})
	cmds.Set("delete", func(name string) {
		for i, cmd := range graph.Commands {
			if cmd.Name == name {
				graph.Commands = slices.Delete(graph.Commands, i, i+1)
				return
			}
		}
	})
	cmds.Set("get", func(name string) *commands.Command {
		for _, cmd := range graph.Commands {
			if cmd.Name == name {
				return cmd
			}
		}
		return nil
	})
	cmds.Set("getAll", func(name string) []*commands.Command {
		return graph.Commands
	})

	server.Set("commands", cmds)

	log.Set("info", func(format string, a ...interface{}) {
		logger.Info(format, a...)
	})

	log.Set("warn", func(format string, a ...interface{}) {
		logger.Warn(format, a...)
	})

	log.Set("debug", func(format string, a ...interface{}) {
		logger.Debug(format, a...)
	})

	log.Set("error", func(format string, a ...interface{}) {
		logger.Error(format, a...)
	})

	log.Set("print", func(format string, a ...interface{}) {
		logger.Print(format, a...)
	})

	vm.Set("Plugin", func(data *pluginConfiguration) {
		if data == nil {
			logger.Error("Failed to load plugin %s: invalid plugin data", plugin.Filename)
		} else {
			if data.Identifier == "" {
				logger.Error("Failed to load plugin %s: identifier was not specified", plugin.Filename)
			}
			plugin.Identifier = data.Identifier
			plugin.OnLoad = data.OnLoad

			data.OnLoad()

			plugin.Initialized = true
		}
	})
	server.Set("logger", log)
	vm.Set("server", server)

	vm.Set("require", func(file string) map[string]interface{} {
		exports := make(map[string]interface{})
		path := root + "/" + file
		f, err := os.ReadFile(path)
		if err != nil {
			return exports
		}
		v := *vm
		v.Set("exports", goja.Undefined())
		v.RunString(string(f))
		v.ExportTo(v.Get("exports"), &exports)
		return exports
	})
	return vm
}
