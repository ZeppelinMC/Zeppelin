package handlers

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/player"
)

func ClientSettings(controller controller, state *player.Player, pk *packet.ClientSettings) {
	old := state.ClientSettings()
	state.SetClientSettings(player.ClientInformation(*pk))
	if old.ViewDistance != pk.ViewDistance {
		controller.SendSpawnChunks()
	}
}
