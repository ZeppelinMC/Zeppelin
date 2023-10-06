package handlers

func PlayerCommand(controller Controller, action int32) {
	switch action {
	case 0:
		{
			controller.BroadcastPose(5)
		}
	case 1:
		{
			controller.BroadcastPose(0)
		}
	case 3:
		{
			controller.BroadcastSprinting(true)
		}
	case 4:
		{
			controller.BroadcastSprinting(false)
		}
	}
}
