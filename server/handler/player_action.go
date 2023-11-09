package handler

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/enum"
	"github.com/dynamitemc/dynamite/server/player"
)

func PlayerAction(state *player.Player, pk *packet.PlayerActionServer) {
	switch pk.Status {
	case enum.PlayerActionStartedDigging:
		state.BroadcastMetadataInArea(&packet.SetEntityMetadata{
			EntityID: state.EntityID(),
			Pose:     &dg,
		})
		//go controller.BroadcastDigging(pk.Location)
	case enum.PlayerActionCancelledDigging, enum.PlayerActionFinishedDigging:
		if pk.Status == 2 {
			state.BreakBlock(pk.Location)
		}
		state.BroadcastMetadataInArea(&packet.SetEntityMetadata{
			EntityID: state.EntityID(),
			Pose:     &st,
		})
	case enum.PlayerActionDropItemStack, enum.PlayerActionDropItem:
		if s, ok := state.Inventory.HeldItem(); ok {
			state.SetPreviousSelectedSlot(s)
			state.Inventory.DeleteSlot(int8(s.Slot))
		}
		//controller.DropSlot()
	}
}
