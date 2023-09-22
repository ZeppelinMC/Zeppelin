package server

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/dynamitemc/dynamite/logger"
	"github.com/robertkrimen/otto"
)

var unrecognizedScript = errors.New("unrecognized scripting language")

type loader struct {
	js *otto.Otto
}

func getJavaScriptVM(logger logger.Logger, pluginName string) *otto.Otto {
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
				logger.Error("Failed to load plugin %s: identifier was not specified", pluginName)
			}

		} else {
			logger.Error("Failed to load plugin %s: invalid plugin data", pluginName)
		}
		return otto.UndefinedValue()
	})
	return vm
}

func LoadPlugins(logger logger.Logger) error {
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
		if err := LoadPlugin("plugins/"+file.Name(), logger, loader{js: getJavaScriptVM(logger, file.Name())}); err != nil {
			return err
		}
	}
	return nil
}

func LoadPlugin(path string, logger logger.Logger, loader loader) error {
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	sp := strings.Split(path, "/")
	filename := sp[len(sp)-1]
	switch {
	case strings.HasSuffix(path, ".js"):
		{
			_, err := loader.js.Run(string(file))
			if err != nil {
				logger.Error("Failed to load plugin %s: %s", filename, err)
			}
		}
	}
	return nil
}
