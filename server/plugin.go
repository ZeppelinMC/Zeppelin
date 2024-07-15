package server

import (
	"plugin"

	"github.com/dynamitemc/aether/log"
)

type Plugin interface {
	Name() string
	OnStart(*Server)
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
	plugin.OnStart(srv)
}
