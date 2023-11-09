package enum

const (
	PoseStanding int32 = iota
	PoseFallFlying
	PoseSleeping
	PoseSwimming
	PoseSpinAttack
	PoseSneaking
	PoseLongJumping
	PoseDying
	PoseCroaking
	PoseUsingTongue
	PoseSitting
	PoseRoaring
	PoseSniffing
	PoseEmerging
	PoseDigging
)

const (
	SnifferStateIdling int32 = iota
	SnifferStateFeelingHappy
	SnifferStateScenting
	SnifferStateSniffing
	SnifferStateSearching
	SnifferStateDigging
	SnifferStateRising
)

const (
	VillagerTypeDesert int32 = iota
	VillagerTypeJungle
	VillagerTypePlains
	VillagerTypeSavanna
	VillagerTypeSnow
	VillagerTypeSwamp
	VillagerTypeTaiga
)

const (
	VillagerProfessionNone int32 = iota
	VillagerProfessionArmorer
	VillagerProfessionButcher
	VillagerProfessionCartographer
	VillagerProfessionCleric
	VillagerProfessionFarmer
	VillagerProfessionFisherman
	VillagerProfessionFletcher
	VillagerProfessionLeatherworker
	VillagerProfessionLibrarian
	VillagerProfessionMason
	VillagerProfessionNitwit
	VillagerProfessionShepherd
	VillagerProfessionToolsmith
	VillagerProfessionWeaponsmith
)

const (
	DirectionDown int32 = iota
	DirectionUp
	DirectionNorth
	DirectionSouth
	DirectionWest
	DirectionEast
)

const (
	EntityMetadataTypeByte int32 = iota
	EntityMetadataTypeVarInt
	EntityMetadataTypeVarLong
	EntityMetadataTypeFloat
	EntityMetadataTypeString
	EntityMetadataTypeChat
	EntityMetadataTypeOptChat
	EntityMetadataTypeSlot
	EntityMetadataTypeBoolean
	EntityMetadataTypeRotation
	EntityMetadataTypePosition
	EntityMetadataTypeOptPosition
	EntityMetadataTypeDirection
	EntityMetadataTypeOptUUID
	EntityMetadataTypeBlockID
	EntityMetadataTypeOptBlockID
	EntityMetadataTypeNBT
	EntityMetadataTypeParticle
	EntityMetadataTypeVillagerData
	EntityMetadataTypeOptVarInt
	EntityMetadataTypePose
	EntityMetadataTypeCatVariant
	EntityMetadataTypeFrogVariant
	EntityMetadataTypeOptGlobalPos
	EntityMetadataTypePaintingVariant
	EntityMetadataTypeSnifferState
	EntityMetadataTypeVector3
	EntityMetadataTypeQuaternion
)
