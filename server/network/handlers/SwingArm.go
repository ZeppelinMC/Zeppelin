package handlers

func SwingArm(controller controller, hand int32) {
	var animation uint8
	if hand == 1 {
		animation = 3
	}
	controller.BroadcastAnimation(animation)
}
