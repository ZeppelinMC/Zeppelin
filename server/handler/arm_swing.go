package handler

import (
	"github.com/dynamitemc/dynamite/server/enum"
	"github.com/dynamitemc/dynamite/server/player"
)

func SwingArm(state *player.Player, hand int32) {
	var animation uint8
	if hand == 1 {
		animation = enum.EntityAnimationSwingOffhand
	}
	state.BroadcastAnimation(animation)
}
