package plugins

import (
	"errors"

	lua "github.com/Shopify/go-lua"
	"github.com/dop251/goja"
)

var PluginUnrecognizedType = errors.New("unrecognized plugin type")
var PluginUninitalized = errors.New("plugin was not initialized")
var PluginNoData = errors.New("no plugin data found")
var PluginInvalidData = errors.New("failed to parse plugin data")
var PluginNoRoot = errors.New("no plugin root file")
var PluginSingleFile = errors.New("cannot import files in single file plugin")

const (
	PluginTypeJavaScript = iota
	PluginTypeLua
)

type PluginData struct {
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
