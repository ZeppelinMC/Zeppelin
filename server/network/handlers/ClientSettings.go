package handlers

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/player"
)

func ClientSettings(controller Controller, state *player.Player, pk *packet.ClientSettings) {
	controller.BroadcastSkinData()
	state.SetClientSettings(player.ClientInformation(*pk))
}
