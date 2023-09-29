package server

import (
	"errors"
	"fmt"
	"io/fs"
	"net/rpc"
	"os"
	"os/exec"

	"github.com/dynamitemc/dynamite/util"
	"github.com/hashicorp/go-plugin"
)

var pluginsDir = util.GetArg("pluginspath", "plugins")

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion: 1,
}

type Plugin struct {
	Identifier string
	OnLoad     func(*Server)
}

func (p *Plugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return p, nil
}

func (p Plugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return p, nil
}

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
			//srv.Plugins[plugin.Identifier] = plugin
			//srv.Logger.Debug("Finished loading plugin %s", plugin.Identifier)
		}
	}
	return nil
}

func (srv *Server) LoadPlugin(path string) (interface{}, error) {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command(path),
	})
	defer client.Kill()

	rpcClient, _ := client.Client()

	p, e := rpcClient.Dispense("plugin")
	fmt.Println(e, p.(*Plugin).Identifier)

	return nil, nil
}

var pluginMap = map[string]plugin.Plugin{
	"plugin": &Plugin{},
}
