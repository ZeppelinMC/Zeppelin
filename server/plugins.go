package server

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"plugin"
	"runtime"
)

type Plugin struct {
	Identifier string
	OnLoad     func(*Server)
}

func (srv *Server) LoadPlugins() {
	if runtime.GOOS != "darwin" && runtime.GOOS == "linux" && runtime.GOOS != "freebsd" {
		srv.Logger.Error("Plugins are not supported on your platform yet. Come back tomorrow.")
		return
	}
	err := os.Mkdir("plugins", 0755)
	if err != nil {
		if !errors.Is(err, fs.ErrExist) {
			srv.Logger.Error("Failed to load plugins.")
			return
		}
	}
	dir, err := os.ReadDir("plugins")
	if err != nil {
		srv.Logger.Error("Failed to load plugins.")
		return
	}
	for _, file := range dir {
		srv.Logger.Debug("Loading plugin %s", file.Name())
		if plugin, err := srv.LoadPlugin("plugins/" + file.Name()); err != nil {
			srv.Logger.Error("Failed to load plugin %s: %s", file.Name(), err)
			return
		} else {
			fmt.Println(plugin)
			srv.Plugins[plugin.Identifier] = plugin
			srv.Logger.Debug("Finished loading plugin %s", plugin.Identifier)
		}
	}
}

func (srv *Server) LoadPlugin(path string) (*Plugin, error) {
	p, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}
	data, err := p.Lookup("Plugin")
	if err != nil {
		return nil, err
	}
	pl, ok := data.(*Plugin)
	if !ok {
		return nil, ErrPluginCantGetData
	}
	return pl, nil
}

var ErrPluginCantGetData = errors.New("couldn't get plugin data")
