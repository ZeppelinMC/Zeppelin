package region

import "github.com/zeppelinmc/zeppelin/server/world/region/block"

type Heightmaps struct {
	MOTION_BLOCKING, MOTION_BLOCKING_NO_LEAVES, OCEAN_FLOOR, WORLD_SURFACE [37]int64
}

type anvilChunk struct {
	DataVersion   int32
	Heightmaps    Heightmaps
	InhabitedTime int64
	LastUpdate    int64
	Status        string
	BlockEntities []struct {
		Id string
		X  int32 `nbt:"x"`
		Y  int32 `nbt:"y"`
		Z  int32 `nbt:"z"`
	} `nbt:"block_entities"`

	Sections []struct {
		BlockLight, SkyLight []byte
		Y                    int8
		Biomes               struct {
			Data    []int64  `nbt:"data"`
			Palette []string `nbt:"palette"`
		} `nbt:"biomes"`
		BlockStates struct {
			Data    []int64       `nbt:"data"`
			Palette []block.Block `nbt:"palette"`
		} `nbt:"block_states"`
	} `nbt:"sections"`

	XPos int32 `nbt:"xPos"`
	YPos int32 `nbt:"yPos"`
	ZPos int32 `nbt:"zPos"`
}
