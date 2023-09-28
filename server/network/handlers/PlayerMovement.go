package handlers

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/player"
)

func PlayerMovement(
	controller controller,
	state *player.Player,
	pk packet.Packet,
) {
	switch pk := pk.(type) {
	case *packet.PlayerPosition:
		{
			state.OnGround = pk.OnGround
			controller.BroadcastMovement(state.X, state.Y, state.Z)
			state.X, state.Y, state.Z = pk.X, pk.FeetY, pk.Z
		}
	case *packet.PlayerPositionRotation:
		{
			state.X, state.Y, state.Z, state.Yaw, state.Pitch, state.OnGround = pk.X, pk.FeetY, pk.Z, pk.Yaw, pk.Pitch, pk.OnGround
		}
	case *packet.PlayerRotation:
		{
			state.Yaw, state.Pitch, state.OnGround = pk.Yaw, pk.Pitch, pk.OnGround
		}
	case *packet.PlayerMovement:
		{
			state.OnGround = pk.OnGround
		}
	}

}
