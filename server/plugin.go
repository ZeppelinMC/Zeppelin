package server

import (
	"io/fs"
	"os"
	"plugin"

	"github.com/dynamitemc/aether/log"
)

type Plugin interface {
	Identifier() string
	OnLoad(*Server)
	Unload()
}

func (srv *Server) loadPlugins() {
	fs.WalkDir(os.DirFS("plugins"), ".", func(path string, _ fs.DirEntry, err error) error {
		if path == "." {
			return nil
		}

		if err != nil {
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
	sym, err := pl.Lookup("AetherPluginExport")
	if err != nil {
		log.Errorf("Couldn't find plugin export for %s: %v\n", name, err)
		return
	}
	plugin, ok := sym.(Plugin)
	if !ok {
		log.Errorf("Invalid plugin export for %s\n", name)
		return
	}
	plugin.OnLoad(srv)
}
