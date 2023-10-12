package chunk

type anvilChunk struct {
	Status string

	XPos int32 `nbt:"xPos"`
	YPos int32 `nbt:"yPos"`
	ZPos int32 `nbt:"zPos"`

	Heightmaps struct {
		MotionBlocking []int64 `nbt:"MOTION_BLOCKING"`
		WorldSurface   []int64 `nbt:"WORLD_SURFACE"`
	}

	Sections []struct {
		Y           int8
		BlockStates struct {
			Data    []int64      `nbt:"data"`
			Palette []blockEntry `nbt:"palette"`
		} `nbt:"block_states"`
		Biomes struct {
			Palette []string `nbt:"palette"`
		} `nbt:"biomes"`
		BlockLight []int8
		SkyLight   []int8
	} `nbt:"sections"`

	DataVersion int32
}

type blockEntry struct {
	Properties map[string]string
	Name       string
}

type Entity struct {
	//Attributes            []map[string]interface{} `nbt:"Attributes"`
	IsBaby           int8      `nbt:"IsBaby"`
	AbsorptionAmount float32   `nbt:"AbsorptionAmount"`
	HandDropChances  []float32 `nbt:"HandDropChances"`
	Motion           []float64 `nbt:"Motion"`
	Fire             int16     `nbt:"Fire"`
	CanPickUpLoot    int8      `nbt:"CanPickUpLoot"`
	OnGround         int8      `nbt:"OnGround"`
	//HandItems             []map[string]interface{}   `nbt:"HandItems"`
	Health                float32   `nbt:"Health"`
	Air                   int16     `nbt:"Air"`
	HurtTime              int16     `nbt:"HurtTime"`
	DrownedConversionTime int32     `nbt:"DrownedConversionTime"`
	FallDistance          float32   `nbt:"FallDistance"`
	UUID                  []int32   `nbt:"UUID"`
	ArmorDropChances      []float32 `nbt:"ArmorDropChances"`
	Pos                   []float64 `nbt:"Pos"`
	DeathTime             int16     `nbt:"DeathTime"`
	LeftHanded            int8      `nbt:"LeftHanded"`
	Rotation              []float32 `nbt:"Rotation"`
	HurtByTimestamp       int32     `nbt:"HurtByTimestamp"`
	FallFlying            int8      `nbt:"FallFlying"`
	PortalCooldown        int32     `nbt:"PortalCooldown"`
	InWaterTime           int32     `nbt:"InWaterTime"`
	Id                    string    `nbt:"id"`
	Brain                 struct {
		Memories struct {
		} `nbt:"memories"`
	} `nbt:"Brain"`
	Invulnerable        int8 `nbt:"Invulnerable"`
	PersistenceRequired int8 `nbt:"PersistenceRequired"`
	CanBreakDoors       int8 `nbt:"CanBreakDoors"`
	//ArmorItems         []map[string]interface{} `nbt:"ArmorItems"`
}
