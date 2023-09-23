package server

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	lua "github.com/Shopify/go-lua"
	"github.com/dop251/goja"
)

var unrecognizedScript = errors.New("unrecognized scripting language")
var uninitalized = errors.New("plugin was not initialized")

const (
	PluginTypeJavaScript = iota
	PluginTypeLua
)

type Plugin struct {
	Identifier  string
	Initialized bool
	Filename    string
	JSLoader    *goja.Runtime
	LuaLoader   *lua.State
	Type        int
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
		srv.Logger.Debug("Loading plugin %s", file.Name())
		if plugin, err := srv.LoadPlugin("plugins/" + file.Name()); err != nil {
			return err
		} else {
			if plugin == nil {
				continue
			}
			srv.Plugins[plugin.Identifier] = plugin
			srv.Logger.Debug("Finished loading plugin %s", plugin.Identifier)
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
			_, err := js.RunString(string(file))
			if err != nil {
				srv.Logger.Error("Failed to load plugin %s: %s", filename, err)
				return nil, err
			}
			if !plugin.Initialized {
				srv.Logger.Error("Failed to load plugin %s: %s", filename, uninitalized)
				return nil, uninitalized
			}
			return plugin, nil
		}
	case strings.HasSuffix(path, ".lua"):
		{
			plugin := &Plugin{Filename: filename, Type: PluginTypeLua}
			l := getLuaVM(srv.Logger, plugin)
			plugin.LuaLoader = l
			lua.DoString(l, string(file))
			if !plugin.Initialized {
				srv.Logger.Error("Failed to load plugin %s: %s", filename, uninitalized)
				return nil, uninitalized
			}
			return plugin, nil
		}
	}
	return nil, nil
}
