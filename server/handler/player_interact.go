package handler

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/player"
)

func Interact(state *player.Player, pk *packet.InteractServer) {
	if pk.Type == 1 {
		state.Attack(pk.EntityID)
	}
}
