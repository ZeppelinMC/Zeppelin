package plugin

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server"
	"github.com/dynamitemc/dynamite/server/player"
)

type Plugin interface {
	// Server returns the server for this plugin
	Server() *server.Server
	// Identifier returns the unique plugin name
	Identifier() string
	// OnLoad gets called when the plugin is being loaded
	OnLoad()
	// OnEnable gets called when the plugin is enabled
	OnEnable()
	// OnDisable gets called when the plugin is disabled
	OnDisable()
}

type PluginPlayerJoin interface {
	Plugin
	// OnPlayerJoin gets called when a player joins the server
	OnPlayerJoin(*player.Player)
}

type PluginPlayerLeave interface {
	Plugin
	// OnPlayerLeave gets called when a player leaves the server
	OnPlayerLeave(*player.Player)
}

type PluginPacketSend interface {
	Plugin
	// OnPacketSend gets called when a packet is sent by the server
	OnPacketSend(packet.Packet, *player.Player)
}

type PluginPacketReceive interface {
	Plugin
	// OnPacketRecieve gets called when a packet is sent to the server by a player
	OnPacketReceive(packet.Packet, *player.Player)
}
