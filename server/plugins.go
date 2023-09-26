package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/Shopify/go-lua"
	"github.com/dynamitemc/dynamite/server/plugins"
	"github.com/dynamitemc/dynamite/util"
)

var pluginsDir = util.GetArg("pluginspath", "plugins")

func (srv *Server) LoadPlugins() error {
	err := os.Mkdir(pluginsDir, 0755)
	if err != nil {
		if !errors.Is(err, fs.ErrExist) {
			return err
		}
	}
	dir, err := os.ReadDir(pluginsDir)
	if err != nil {
		fmt.Println(err)
		return err
	}
	for _, file := range dir {
		srv.Logger.Debug("Loading plugin %s", file.Name())
		if plugin, err := srv.LoadPlugin(pluginsDir + "/" + file.Name()); err != nil {
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

func (srv *Server) LoadPlugin(path string) (*plugins.Plugin, error) {
	info, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if info.IsDir() {
		file, err := os.ReadFile(path + "/plugin.json")
		var data plugins.PluginData
		if err != nil {
			srv.Logger.Error("Failed to load plugin %s: %s", info.Name(), plugins.PluginNoData)
			return nil, plugins.PluginNoData
		}
		err = json.Unmarshal(file, &data)
		if err != nil {
			srv.Logger.Error("Failed to load plugin %s: %s", info.Name(), plugins.PluginInvalidData)
			return nil, plugins.PluginInvalidData
		}
		file, err = os.ReadFile(path + "/" + data.RootFile)
		if err != nil {
			srv.Logger.Error("Failed to load plugin %s: %s", info.Name(), plugins.PluginNoRoot)
			return nil, plugins.PluginNoRoot
		}
		switch data.Type {
		case "javascript":
			{
				plugin := &plugins.Plugin{Filename: info.Name(), Type: plugins.PluginTypeJavaScript}
				js := plugins.GetJavaScriptVM(srv.Logger, plugin, path)
				plugin.JSLoader = js
				_, err := js.RunString(string(file))
				if err != nil {
					srv.Logger.Error("Failed to load plugin %s: %s", info.Name(), err)
					return nil, err
				}
				if !plugin.Initialized {
					srv.Logger.Error("Failed to load plugin %s: %s", info.Name(), plugins.PluginUninitalized)
					return nil, plugins.PluginUninitalized
				}
				return plugin, nil
			}
		case "lua":
			{
				plugin := &plugins.Plugin{Filename: info.Name(), Type: plugins.PluginTypeLua}
				l := plugins.GetLuaVM(srv.Logger, plugin)
				plugin.LuaLoader = l
				lua.DoString(l, string(file))
				if !plugin.Initialized {
					srv.Logger.Error("Failed to load plugin %s: %s", info.Name(), plugins.PluginUninitalized)
					return nil, plugins.PluginUninitalized
				}
				return plugin, nil
			}
		default:
			{
				srv.Logger.Error("Failed to load plugin %s: %s", info.Name(), plugins.PluginUnrecognizedType)
				return nil, plugins.PluginUnrecognizedType
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
				plugin := &plugins.Plugin{Filename: filename, Type: plugins.PluginTypeJavaScript}
				js := plugins.GetJavaScriptVM(srv.Logger, plugin, "")
				plugin.JSLoader = js
				_, err := js.RunString(string(file))
				if err != nil {
					srv.Logger.Error("Failed to load plugin %s: %s", filename, err)
					return nil, err
				}
				if !plugin.Initialized {
					srv.Logger.Error("Failed to load plugin %s: %s", filename, plugins.PluginUninitalized)
					return nil, plugins.PluginUninitalized
				}
				return plugin, nil
			}
		case strings.HasSuffix(path, ".lua"):
			{
				plugin := &plugins.Plugin{Filename: filename, Type: plugins.PluginTypeLua}
				l := plugins.GetLuaVM(srv.Logger, plugin)
				plugin.LuaLoader = l
				lua.DoString(l, string(file))
				if !plugin.Initialized {
					srv.Logger.Error("Failed to load plugin %s: %s", filename, plugins.PluginUninitalized)
					return nil, plugins.PluginUninitalized
				}
				return plugin, nil
			}
		default:
			{
				srv.Logger.Error("Failed to load plugin %s: %s", filename, plugins.PluginUnrecognizedType)
				return nil, plugins.PluginUnrecognizedType
			}
		}
	}
}
