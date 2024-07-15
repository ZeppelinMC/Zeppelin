package server

import (
	"fmt"
	"net/rpc"
	"os/exec"

	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
)

type Plugin interface {
	OnStart() string
}

var PluginHandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "aether",
	MagicCookieValue: "aethermc-da serber",
}

type pluginRPCServer struct {
	impl Plugin
}

func (s *pluginRPCServer) OnStart(args interface{}, resp *string) error {
	fmt.Println("pluginrpcserver onstart")
	*resp = s.impl.OnStart()
	return nil
}

type pluginRPC struct{ client *rpc.Client }

func (g *pluginRPC) OnStart() string {
	var resp string
	err := g.client.Call("Plugin.OnStart", new(interface{}), &resp)
	fmt.Println(err)
	if err != nil {
		// You usually want your interfaces to return errors. If they don't,
		// there isn't much other choice here.
		panic(err)
	}
	return resp
}

type PlugImpl struct {
	Impl Plugin
}

func (p *PlugImpl) Server(*plugin.MuxBroker) (interface{}, error) {
	return &pluginRPCServer{impl: p.Impl}, nil
}

func (PlugImpl) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &pluginRPC{client: c}, nil
}

func (srv *Server) LoadPlugin() {
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig: PluginHandshakeConfig,
		Plugins:         plugin.PluginSet{"plugin": &PlugImpl{}},
		Cmd:             exec.Command("./plugin.exe"),
		Logger:          hclog.New(&hclog.LoggerOptions{Level: 6}),
	})
	defer client.Kill()

	rpcClient, err := client.Client()
	if err != nil {
		fmt.Println(err)
		return
	}

	raw, err := rpcClient.Dispense("plugin")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T\n", raw)
	fmt.Println(raw.(Plugin).OnStart())
}
