package handlers

import "github.com/dynamitemc/dynamite/server/player"

func SetHeldItem(state *player.Player, heldItem int16) {
	state.SetHeldItem(int32(heldItem))
}
