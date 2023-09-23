package server

import (
	"os"

	"github.com/dop251/goja"
	"github.com/dynamitemc/dynamite/logger"
)

type PluginData struct {
	Identifier string `js:"identifier"`
}

func makeJSMap[K comparable, V any](vm *goja.Runtime, from *map[K]V) *goja.Object {
	m := vm.NewObject()
	m.Set("set", func(key K, value V) {
		(*from)[key] = value
	})
	m.Set("delete", func(key K) {
		delete(*from, key)
	})
	m.Set("get", func(key K) V {
		return (*from)[key]
	})
	return m
}

func at[T any](arr []T, index int) (val T) {
	if len(arr) <= index {
		return val
	}
	return arr[index]
}

func getJavaScriptVM(logger logger.Logger, plugin *Plugin) *goja.Runtime {
	vm := goja.New()
	vm.SetFieldNameMapper(goja.TagFieldNameMapper("js", true))
	server := vm.NewObject()
	log := vm.NewObject()

	server.Set("close", func(c int64) {
		var code int64 = c
		os.Exit(int(code))
	})

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

	vm.Set("Plugin", func(data *PluginData) {
		if data == nil {
			logger.Error("Failed to load plugin %s: invalid plugin data", plugin.Filename)
		} else {
			if data.Identifier == "" {
				logger.Error("Failed to load plugin %s: identifier was not specified", plugin.Filename)
			}
			plugin.Identifier = data.Identifier
			plugin.Initialized = true
		}
	})
	server.Set("logger", log)
	vm.Set("server", server)
	return vm
}
