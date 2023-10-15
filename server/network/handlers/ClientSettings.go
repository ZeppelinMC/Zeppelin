package handlers

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/player"
)

func ClientSettings(controller Controller, state *player.Player, pk *packet.ClientSettings) {
	state.SetClientSettings(player.ClientInformation(*pk))
	controller.BroadcastSkinData()
	controller.PlaylistUpdate()
}
