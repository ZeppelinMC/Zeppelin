package handlers

import "github.com/dynamitemc/dynamite/server/player"

func TeleportToEntity(controller Controller, state *player.Player, uuid [16]byte) {
	if state.GameMode() != 3 {
		controller.Disconnect("Yo how do you do dat without gamemode spectator?")
		return
	}
	controller.TeleportToEntity(uuid)
}
