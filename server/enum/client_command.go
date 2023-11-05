package enum

const (
	ClientCommandRespawn int32 = iota
	ClientCommandRequestStats
)

const (
	PlayerCommandStartSneaking int32 = iota
	PlayerCommandStopSneaking
	PlayerCommandLeaveBed
	PlayerCommandStartSprinting
	PlayerCommandStopSprinting
	PlayerCommandStartJumpWithHorse
	PlayerCommandStopJumpWithHorse
	PlayerCommandOpenHorseInventory
	PlayerCommandStartFlyingElytra
)

const (
	PlayerActionStartedDigging int32 = iota
	PlayerActionCancelledDigging
	PlayerActionFinishedDigging
	PlayerActionDropItemStack
	PlayerActionDropItem
	PlayerActionShootArrowFinishEating
	PlayerActionSwapItemInHand
)
