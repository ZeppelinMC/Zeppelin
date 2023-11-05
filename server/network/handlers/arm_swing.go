package handlers

import "github.com/dynamitemc/dynamite/server/enum"

func SwingArm(controller Controller, hand int32) {
	var animation uint8
	if hand == 1 {
		animation = enum.EntityAnimationSwingOffhand
	}
	controller.BroadcastAnimation(animation)
}
