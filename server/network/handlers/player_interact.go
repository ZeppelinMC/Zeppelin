package handlers

import "github.com/aimjel/minecraft/packet"

func Interact(controller Controller, pk *packet.InteractServer) {
	if pk.Type == 1 {
		controller.Attack(pk.EntityID)
	}
}
