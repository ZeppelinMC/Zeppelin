package plugins

import (
	"errors"

	lua "github.com/Shopify/go-lua"
	"github.com/dop251/goja"
)

var ErrPluginUnrecognizedType = errors.New("unrecognized plugin type")
var ErrPluginUninitalized = errors.New("plugin was not initialized")
var ErrPluginNoData = errors.New("no plugin data found")
var ErrPluginInvalidData = errors.New("failed to parse plugin data")
var ErrPluginNoRoot = errors.New("no plugin root file")
var ErrPluginSingleFile = errors.New("cannot import files in single file plugin")

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
	OnLoad      func()
}
