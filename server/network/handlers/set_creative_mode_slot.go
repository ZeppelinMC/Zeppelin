package handlers

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/player"
)

func SetCreativeModeSlot(controller Controller, state *player.Player, slot int16, data packet.Slot) {
	if state.GameMode() != 1 {
		controller.Disconnect("bruh cant use de creative button without creative")
		return
	}
	if !data.Present {
		if s, ok := state.InventorySlot(int(slot)); ok {
			state.SetPreviousSelectedSlot(s)
		}
		controller.ClearItem(slot)
	} else {
		controller.SetSlot(slot, data)
	}
}
