package handlers

import (
	"github.com/aimjel/minecraft/packet"
	"github.com/dynamitemc/dynamite/server/player"
)

func PlayerMovement(
	state *player.Player,
	pk packet.Packet,
) {
	switch pk := pk.(type) {
	case *packet.PlayerPosition:
		{
			state.X, state.Y, state.Z, state.OnGround = pk.X, pk.FeetY, pk.Z, pk.OnGround
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
