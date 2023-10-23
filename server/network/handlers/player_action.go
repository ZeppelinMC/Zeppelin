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
			/*pos := int64(pk.Location)
			x := pos >> 38
			y := pos << 52 >> 52
			z := pos << 26 >> 38
			fmt.Println("broke", x, y, z)*/
			controller.BreakBlock(pk.Location)
		}
		controller.BroadcastPose(0)
	case 3, 4:
		if s, ok := state.InventorySlot(int(state.HeldItem())); ok {
			state.SetPreviousSelectedSlot(s)
			state.DeleteInventorySlot(int(state.HeldItem()))
		}
		controller.DropSlot()
	}
}
