package handlers

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/player"
)

func PlayerAction(controller Controller, state *player.Player, pk *packet.PlayerActionServer) {
	switch pk.Status {
	case 0:
		controller.BroadcastPose(14)
		//go controller.BroadcastDigging(pk.Location)
	case 1, 2:
		if pk.Status == 2 {
			controller.BreakBlock(pk.Location)
		}
		controller.BroadcastPose(0)
	case 3, 4:
		if s, ok := state.Inventory().Slot(int8(state.HeldItem())); ok {
			state.SetPreviousSelectedSlot(s)
			state.Inventory().DeleteSlot(int8(state.HeldItem()))
		}
		controller.DropSlot()
	}
}
