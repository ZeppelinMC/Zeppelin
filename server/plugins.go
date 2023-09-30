package server

import (
	"errors"
	"fmt"
	"io/fs"
	"net/rpc"
	"os"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

var handshakeConfig = plugin.HandshakeConfig{
	MagicCookieKey:   "Plugin",
	MagicCookieValue: "Plugin",
	ProtocolVersion:  1,
}

type Plugin struct {
	Identifier string
	OnLoad     func(*Server)
}

func (p *Plugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return p, nil
}

func (p *Plugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return p, nil
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
		Logger: hclog.New(&hclog.LoggerOptions{
			Level: -1,
		}),
	})
	defer client.Kill()

	rpcClient, _ := client.Client()

	p, err := rpcClient.Dispense("plugin")
	fmt.Println("yeah!", p.(*Plugin).Identifier, err)

	return nil, nil
}

var pluginMap = map[string]plugin.Plugin{
	"plugin": &Plugin{},
}
