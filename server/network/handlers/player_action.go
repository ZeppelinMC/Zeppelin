package handlers

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/enum"
	"github.com/dynamitemc/dynamite/server/player"
)

func PlayerAction(controller Controller, state *player.Player, pk *packet.PlayerActionServer) {
	switch pk.Status {
	case enum.PlayerActionStartedDigging:
		controller.BroadcastPose(14)
		//go controller.BroadcastDigging(pk.Location)
	case enum.PlayerActionCancelledDigging, enum.PlayerActionFinishedDigging:
		if pk.Status == 2 {
			controller.BreakBlock(pk.Location)
		}
		controller.BroadcastPose(0)
	case enum.PlayerActionDropItemStack, enum.PlayerActionDropItem:
		if s, ok := state.Inventory().HeldItem(); ok {
			state.SetPreviousSelectedSlot(s)
			state.Inventory().DeleteSlot(int8(s.Slot))
		}
		controller.DropSlot()
	}
}
