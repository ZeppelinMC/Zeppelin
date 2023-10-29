package handlers

import "github.com/dynamitemc/dynamite/server/player"

func SetHeldItem(state *player.Player, heldItem int16) {
	state.SetSelectedSlot(int32(heldItem))
}
