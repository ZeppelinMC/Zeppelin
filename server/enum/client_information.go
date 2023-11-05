package enum

const (
	ChatModeEnabled int32 = iota
	ChatModeCommandsOnly
	ChatModeHidden
)

const (
	MainHandLeft int32 = iota
	MainHandRight
)

const (
	DisplayedSkinPartCape        byte = 0x01
	DisplayedSkinPartJacket      byte = 0x02
	DisplayedSkinPartLeftSleeve  byte = 0x04
	DisplayedSkinPartRightSleeve byte = 0x08
	DisplayedSkinPartLeftLeg     byte = 0x10
	DisplayedSkinPartRightLeg    byte = 0x20
	DisplayedSkinPartHat         byte = 0x40
)
