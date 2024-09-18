package server

import (
	"io/fs"
	"os"
	"plugin"

	"github.com/zeppelinmc/zeppelin/util/log"
)

type Plugin struct {
	basePluginsPath string

	srv        *Server
	Identifier string

	OnLoad func(*Plugin)
	Unload func(*Plugin)
}

func (p Plugin) FS() fs.FS {
	return os.DirFS(p.Dir())
}

// Dir returns the base directory for the plugin (plugins/<identifier>)
func (p Plugin) Dir() string {
	return p.basePluginsPath + "/" + p.Identifier
}

func (p Plugin) Server() *Server {
	return p.srv
}

func (srv *Server) loadPlugins() {
	os.Mkdir("plugins", 0755)
	dir, _ := os.ReadDir("plugins")
	for _, entry := range dir {
		if entry.IsDir() {
			continue
		}
		srv.loadPlugin("plugins/" + entry.Name())
	}
}

func (srv *Server) loadPlugin(name string) {
	pl, err := plugin.Open(name)
	if err != nil {
		log.Errorlnf("Error loading plugin %s: %v", name, err)
		return
	}
	sym, err := pl.Lookup("ZeppelinPluginExport")
	if err != nil {
		log.Errorlnf("Couldn't find plugin export for %s: %v", name, err)
		return
	}
	plugin, ok := sym.(*Plugin)
	if !ok {
		log.Errorlnf("Invalid plugin export for %s", name)
		return
	}
	plugin.basePluginsPath = "plugins"
	plugin.srv = srv
	plugin.OnLoad(plugin)
	os.Mkdir(plugin.Dir(), 0755)
}
