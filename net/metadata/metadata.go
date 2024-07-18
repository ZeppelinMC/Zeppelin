package metadata

import (
	"github.com/dynamitemc/aether/chat"
	"github.com/google/uuid"
)

/*
	[index] -> value

value must be one of the types below

learn more about metadata at https://wiki.vg/Entity_metadata
*/
type Metadata map[byte]any

type (
	Byte    int8    //0
	VarInt  int32   //1
	VarLong int64   //2
	Float   float32 //3
	String  string  //4

	TextComponent         chat.TextComponent  //5
	OptionalTextComponent *chat.TextComponent //6
	//Slot //7
	Boolean            bool       //8
	Rotations          [3]Float   //9
	Position           [3]int32   //10
	OptionalPosition   *[3]int32  //11
	Direction          VarInt     //12
	OptionalUUID       *uuid.UUID //13
	BlockState         VarInt     //14
	OptionalBlockState VarInt     //15
	NBT                any        //16
	// Particle //17
	VillagerData   [3]Float //18 | [type, profession, level]
	OptionalVarInt VarInt   //19
	Pose           VarInt   //20
	CatVariant     VarInt   //21
	FrogVariant    VarInt   //22

	GlobalPosition struct {
		DimensionIdentifier String
		Position            Position
	}
	OptionalGlobalPosition *GlobalPosition //23

	PaintingVariant VarInt   //24
	SnifferState    VarInt   //25
	Vector3         [3]Float //26
	Quatermion      [4]Float //27
)

const (
	Standing Pose = iota
	FallFlying
	Sleeping
	Swimming
	SpinAttack
	Sneaking
	LongJumping
	Dying
	Croaking
	UsingTongue
	Sitting
	Roaring
	Sniffing
	Emerging
	Digging
)

const (
	SnifferIdling SnifferState = iota
	SnifferFeelingHappy
	SnifferScenting
	SnifferSniffing
	SnifferSearching
	SnifferDigging
	SnifferRising
)

const (
	VillagerTypeDesert = iota
	VillagerTypeJungle
	VillagerTypePlains
	VillagerTypeSavanna
	VillagerTypeSnow
	VillagerTypeSwamp
	VillagerTypeTaiga
)

const (
	VillagerProfessionNone = iota
	VillagerProfessionArmorer
	VillagerProfessionButcher
	VillagerProfessionCartographer
	VillagerProfessionCleric
	VillagerProfessionFarmer
	VillagerProfessionFisherman
	VillagerProfessionFletcher
	VillagerProfessionLeatherworker
	VillagerProfessionLibrarian
	VillagerProfessionMason // hi mason
	VillagerProfessionNitwit
	VillagerProfessionShepherd
	VillagerProfessionToolsmith
	VillagerProfessionWeaponsmith
)

const (
	IsOnFire = 1 << iota
	IsCrouching
	IsRiding_unused
	IsSprinting
	IsSwimming
	IsInvisible
	HasGlowingEffect
	IsFlyingWithElytra
)

// base entity
const (
	// Byte (0)
	BaseIndex = iota
	// VarInt (1)
	AirTicksIndex
	// OptionalTextComponent (6)
	CustomNameIndex
	// Boolean (8)
	IsCustomNameVisibleIndex
	// Boolean (8)
	IsSilentIndex
	// Boolean (8)
	HasNoGravityIndex
	// Pose (21)
	PoseIndex
	// VarInt (1)
	TicksFrozenInPowderedSnowIndex
)

// player extends living entity
const (
	// Float (3)
	PlayerAdditionalHeartsIndex = iota + 15
	// VarInt (1)
	PlayerScoreIndex
	// Byte (0)
	PlayerDisplayedSkinPartsIndex
	// Byte (0)
	PlayerMainHandIndex
	// NBT (16)
	LeftShoulderEntityData
	// NBT (16)
	RightShoulderEntityData
)
