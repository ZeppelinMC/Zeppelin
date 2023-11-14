package handler

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/player"
)

func PlayerMovement(
	state *player.Player,
	p packet.Packet,
) {
	if state.IsDead() {
		return
	}
	x, y, z := state.Position()
	yaw, pitch := state.Rotation()
	switch pk := p.(type) {
	case *packet.PlayerPosition:
		{
			state.HandleMovement(pk.ID(), pk.X, pk.FeetY, pk.Z, yaw, pitch, pk.OnGround, false)
		}
	case *packet.PlayerPositionRotation:
		{
			state.HandleMovement(pk.ID(), pk.X, pk.FeetY, pk.Z, pk.Yaw, pk.Pitch, pk.OnGround, false)
		}
	case *packet.PlayerRotation:
		{
			state.HandleMovement(pk.ID(), x, y, z, pk.Yaw, pk.Pitch, pk.OnGround, false)
		}
	}
}
