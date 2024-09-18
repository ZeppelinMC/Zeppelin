// Package region provides decoding and encoding of Region (.mca) files

package region

import (
	"unsafe"

	"github.com/zeppelinmc/zeppelin/server/world/chunk"
	"github.com/zeppelinmc/zeppelin/server/world/chunk/heightmaps"
)

const DataVersion = 3953

// chunkToAnvil turns the chunk into an anvilChunk
func chunkToAnvil(c *chunk.Chunk) anvilChunk {
	anvil := anvilChunk{
		Status:        "minecraft:full",
		BlockEntities: c.BlockEntities,
		Heightmaps:    c.Heightmaps,
		Sections:      make([]anvilSection, len(c.Sections)),

		XPos: c.X, YPos: c.Y, ZPos: c.Z,
		DataVersion: DataVersion,
	}

	for i, sec := range c.Sections {
		sky, block := sec.Light()
		_, biomePalette, biomeStates := sec.Biomes()
		_, blockPalette, blockStates := sec.BlockStates()

		anvil.Sections[i] = anvilSection{
			Y:        sec.Y(),
			SkyLight: *(*[]int8)(unsafe.Pointer(&sky)), BlockLight: *(*[]int8)(unsafe.Pointer(&block)),
			Biomes: anvilBiomes{
				Data:    biomeStates,
				Palette: biomePalette,
			},
			BlockStates: anvilBlockStates{
				Data:    blockStates,
				Palette: make([]anvilBlock, len(blockPalette)),
			},
		}

		for bi, block := range blockPalette {
			name, properties := block.Encode()
			anvil.Sections[i].BlockStates.Palette[bi] = anvilBlock{
				Name: name, Properties: properties,
			}
		}
	}

	return anvil
}

type anvilBlock struct {
	Properties map[string]string
	Name       string
}

type anvilSection struct {
	BlockLight  []int8 `nbt:"BlockLight,omitempty"`
	SkyLight    []int8 `nbt:"SkyLight,omitempty"`
	Y           int8
	Biomes      anvilBiomes      `nbt:"biomes"`
	BlockStates anvilBlockStates `nbt:"block_states"`
}

type anvilBiomes struct {
	Data    []int64  `nbt:"data,omitempty"`
	Palette []string `nbt:"palette"`
}

type anvilBlockStates struct {
	Data    []int64      `nbt:"data,omitempty"`
	Palette []anvilBlock `nbt:"palette"`
}

type anvilChunk struct {
	DataVersion   int32
	Heightmaps    heightmaps.Heightmaps
	InhabitedTime int64
	LastUpdate    int64
	Status        string
	BlockEntities []chunk.BlockEntity `nbt:"block_entities,omitempty"`

	Sections []anvilSection `nbt:"sections"`

	XPos int32 `nbt:"xPos"`
	YPos int32 `nbt:"yPos"`
	ZPos int32 `nbt:"zPos"`
}
