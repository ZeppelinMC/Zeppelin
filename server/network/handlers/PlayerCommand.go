package handlers

func PlayerCommand(controller controller, action int32) {
	switch action {
	case 0:
		{
			controller.BroadcastPose(5)
		}
	case 1:
		{
			controller.BroadcastPose(0)
		}
	}
}
