package handlers

import (
	"github.com/aimjel/minecraft/packet"
)

func PlayerAction(controller controller, pk *packet.PlayerActionServer) {
	switch pk.Status {
	case 0:
		controller.BroadcastPose(14)
	case 1, 2:
		controller.BroadcastPose(0)
	}
}
