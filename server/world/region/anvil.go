package region

import (
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
	"github.com/zeppelinmc/zeppelin/server/world/chunk/heightmaps"
)

type anvilBlock struct {
	Properties map[string]string
	Name       string
}

type anvilChunk struct {
	DataVersion   int32
	Heightmaps    heightmaps.Heightmaps
	InhabitedTime int64
	LastUpdate    int64
	Status        string
	BlockEntities []chunk.BlockEntity `nbt:"block_entities"`

	Sections []struct {
		BlockLight, SkyLight []byte
		Y                    int8
		Biomes               struct {
			Data    []int64  `nbt:"data"`
			Palette []string `nbt:"palette"`
		} `nbt:"biomes"`
		BlockStates struct {
			Data    []int64      `nbt:"data"`
			Palette []anvilBlock `nbt:"palette"`
		} `nbt:"block_states"`
	} `nbt:"sections"`

	XPos int32 `nbt:"xPos"`
	YPos int32 `nbt:"yPos"`
	ZPos int32 `nbt:"zPos"`
}
