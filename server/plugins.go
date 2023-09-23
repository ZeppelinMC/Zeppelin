package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	lua "github.com/Shopify/go-lua"
	"github.com/dop251/goja"
)

var pluginUnrecognizedType = errors.New("unrecognized plugin type")
var pluginUninitalized = errors.New("plugin was not initialized")
var pluginNoData = errors.New("no plugin data found")
var pluginInvalidData = errors.New("failed to parse plugin data")
var pluginNoRoot = errors.New("no plugin root file")
var pluginSingleFile = errors.New("cannot import files in single file plugin")

const (
	PluginTypeJavaScript = iota
	PluginTypeLua
)

type pluginData struct {
	RootFile string `json:"rootFile"`
	Type     string `json:"type"`
}

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
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if info.IsDir() {
		file, err := os.ReadFile(path + "/plugin.json")
		var data pluginData
		if err != nil {
			srv.Logger.Error("Failed to load plugin %s: %s", info.Name(), pluginNoData)
			return nil, pluginNoData
		}
		err = json.Unmarshal(file, &data)
		if err != nil {
			srv.Logger.Error("Failed to load plugin %s: %s", info.Name(), pluginInvalidData)
			return nil, pluginInvalidData
		}
		file, err = os.ReadFile(path + "/" + data.RootFile)
		if err != nil {
			srv.Logger.Error("Failed to load plugin %s: %s", info.Name(), pluginNoRoot)
			return nil, pluginNoRoot
		}
		switch data.Type {
		case "javascript":
			{
				plugin := &Plugin{Filename: info.Name(), Type: PluginTypeJavaScript}
				js := getJavaScriptVM(srv.Logger, plugin, path)
				plugin.JSLoader = js
				_, err := js.RunString(string(file))
				if err != nil {
					srv.Logger.Error("Failed to load plugin %s: %s", info.Name(), err)
					return nil, err
				}
				if !plugin.Initialized {
					srv.Logger.Error("Failed to load plugin %s: %s", info.Name(), pluginUninitalized)
					return nil, pluginUninitalized
				}
				return plugin, nil
			}
		case "lua":
			{
				plugin := &Plugin{Filename: info.Name(), Type: PluginTypeLua}
				l := getLuaVM(srv.Logger, plugin)
				plugin.LuaLoader = l
				lua.DoString(l, string(file))
				if !plugin.Initialized {
					srv.Logger.Error("Failed to load plugin %s: %s", info.Name(), pluginUninitalized)
					return nil, pluginUninitalized
				}
				return plugin, nil
			}
		default:
			{
				srv.Logger.Error("Failed to load plugin %s: %s", info.Name(), pluginUnrecognizedType)
				return nil, pluginUnrecognizedType
			}
		}
	} else {
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
				js := getJavaScriptVM(srv.Logger, plugin, "")
				plugin.JSLoader = js
				_, err := js.RunString(string(file))
				if err != nil {
					srv.Logger.Error("Failed to load plugin %s: %s", filename, err)
					return nil, err
				}
				if !plugin.Initialized {
					srv.Logger.Error("Failed to load plugin %s: %s", filename, pluginUninitalized)
					return nil, pluginUninitalized
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
					srv.Logger.Error("Failed to load plugin %s: %s", filename, pluginUninitalized)
					return nil, pluginUninitalized
				}
				return plugin, nil
			}
		default:
			{
				srv.Logger.Error("Failed to load plugin %s: %s", filename, pluginUnrecognizedType)
				return nil, pluginUnrecognizedType
			}
		}
	}
}
