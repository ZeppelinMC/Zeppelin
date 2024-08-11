package server

import (
	"io/fs"
	"os"
	"plugin"

	"github.com/zeppelinmc/zeppelin/log"
)

type Plugin struct {
	srv        *Server
	Identifier string

	OnLoad func(*Plugin)
	Unload func(*Plugin)
}

func (p Plugin) Server() *Server {
	return p.srv
}

func (srv *Server) loadPlugins() {
	os.Mkdir("plugins", 0755)
	fs.WalkDir(os.DirFS("plugins"), ".", func(path string, e fs.DirEntry, err error) error {
		if path == "." || err != nil || e.IsDir() {
			return nil
		}

		srv.loadPlugin("plugins/" + path)
		return nil
	})
}

func (srv *Server) loadPlugin(name string) {
	pl, err := plugin.Open(name)
	if err != nil {
		log.Errorf("Error loading plugin %s: %v\n", name, err)
		return
	}
	sym, err := pl.Lookup("ZeppelinPluginExport")
	if err != nil {
		log.Errorf("Couldn't find plugin export for %s: %v\n", name, err)
		return
	}
	plugin, ok := sym.(*Plugin)
	if !ok {
		log.Errorf("Invalid plugin export for %s\n", name)
		return
	}
	plugin.srv = srv
	plugin.OnLoad(plugin)
}
