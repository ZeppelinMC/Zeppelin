package handlers

import "github.com/dynamitemc/dynamite/server/enum"

func PlayerCommand(controller Controller, action int32) {
	switch action {
	case enum.PlayerCommandStartSneaking: // start sneaking / swimming
		{
			b := controller.OnBlock()
			if b.EncodedName() == "minecraft:water" {
				controller.BroadcastPose(3)
			} else {
				controller.BroadcastPose(5)
			}
		}
	case enum.PlayerCommandStopSneaking: // stop sneaking / swimming
		{
			controller.BroadcastPose(0)
		}
	case enum.PlayerCommandStartSprinting: // sprint
		{
			controller.BroadcastSprinting(true)
		}
	case enum.PlayerCommandStopSprinting: // stop sprint
		{
			controller.BroadcastSprinting(false)
		}
	}
}
