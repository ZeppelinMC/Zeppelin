package handlers

import "github.com/aimjel/minecraft/packet"

func Interact(controller controller, pk *packet.InteractServer) {
	if pk.Type == 1 {
		controller.Hit(pk.EntityID)
	}
}
