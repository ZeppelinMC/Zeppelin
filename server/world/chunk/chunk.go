// Package chunk provides a way of encoding and modifying chunks
package chunk

import (
	"fmt"
	"unsafe"

	"github.com/zeppelinmc/zeppelin/server/container"
	"github.com/zeppelinmc/zeppelin/server/world/chunk/heightmaps"
	"github.com/zeppelinmc/zeppelin/server/world/chunk/section"
)

type Generator interface {
	NewChunk(x, z int32) Chunk
	GenerateWorldSpawn() (x, y, z int32)
}

const MinChunkY = -4

type BlockEntity struct {
	Id         string `nbt:"id"`
	X          int32  `nbt:"x"`
	Y          int32  `nbt:"y"`
	Z          int32  `nbt:"z"`
	KeepPacked bool   `nbt:"keepPacked"`

	// for chest entities
	Items container.Container `nbt:"Items,omitempty"`
}

type Chunk struct /*{
	//LastModified int64
	X, Y, Z    int32
	Heightmaps heightmaps.Heightmaps

	Sections      []*section.Section
	BlockEntities []BlockEntity
}*/{
	DataVersion   int32
	Heightmaps    heightmaps.Heightmaps
	InhabitedTime int64
	LastUpdate    int64
	Status        string
	BlockEntities []BlockEntity `nbt:"block_entities,omitempty"`

	Sections []section.Section `nbt:"sections"`

	X int32 `nbt:"xPos"`
	Y int32 `nbt:"yPos"`
	Z int32 `nbt:"zPos"`
}

func NewChunk(x, z int32) Chunk {
	c := Chunk{
		Y: MinChunkY,
		X: x,
		Z: z,
		//LastModified: time.Now().UnixMilli(),

		Sections: make([]section.Section, 24),
	}

	for i := range c.Sections {
		c.Sections[i] = section.Section{
			Y: int8(i - MinChunkY),
			BlockStates: section.AnvilBlockStates{
				Palette: []section.AnvilBlock{{Name: "minecraft:air"}},
			},
			Biomes: section.AnvilBiomes{
				Palette: []string{"minecraft:plains"},
			},
			SkyLight: *(*[]int8)(unsafe.Pointer(&fullLightBuffer)),
		}
	}

	return c
}

// X and Z should be relative to the chunk (aka x&0x0f, z&0x0f), but Y should be absolute.
func (chunk *Chunk) Block(x, y, z int32) (section.AnvilBlock, error) {
	secIndex := (y >> 4) - chunk.Y
	if secIndex < 0 || secIndex >= int32(len(chunk.Sections)) {
		return section.AnvilBlock{}, fmt.Errorf("null section")
	}
	sec := chunk.Sections[secIndex]

	return sec.Block(byte(x), byte(y)&0x0f, byte(z)), nil
}

// This function does not update the block for the players, so it should not be used. X and Z should be relative to the chunk (aka x&0x0f, z&0x0f), but Y should be absolute.
func (chunk *Chunk) SetBlock(x, y, z int32, b section.AnvilBlock) (state int64, err error) {
	//c.LastModified = time.Now().UnixMilli()
	secIndex := (y >> 4) - chunk.Y
	if secIndex < 0 || secIndex >= int32(len(chunk.Sections)) {
		return 0, fmt.Errorf("null section")
	}
	sec := chunk.Sections[secIndex]
	return sec.SetBlock(byte(x), byte(y)&0x0f, byte(z), b), nil
}

// This function does not update the block for the players, so it should not be used. X and Z should be relative to the chunk (aka x&0x0f, z&0x0f), but Y should be absolute.
func (chunk *Chunk) SetSkylightLevel(x, y, z int32, value byte) error {
	//c.LastModified = time.Now().UnixMilli()
	secIndex := (y >> 4) - chunk.Y
	if secIndex < 0 || secIndex >= int32(len(chunk.Sections)) {
		return fmt.Errorf("null section")
	}
	sec := chunk.Sections[secIndex]
	return sec.SetSkylightLevel(int(x), int(y)&0x0f, int(z), value)
}

// X and Z should be relative to the chunk (aka x&0x0f, z&0x0f), but Y should be absolute.
func (chunk *Chunk) SkyLightLevel(x, y, z int32) (byte, error) {
	secIndex := (y >> 4) - chunk.Y
	if secIndex < 0 || secIndex >= int32(len(chunk.Sections)) {
		return 0, fmt.Errorf("null section")
	}
	sec := chunk.Sections[secIndex]
	return sec.SkyLightLevel(int(x), int(y)&0x0f, int(z))
}

// Returns the block state at the position. All of the position values should be absolute (aka (chunkPos<<4)+pos)
func (chunk *Chunk) BlockEntity(x, y, z int32) (*BlockEntity, bool) {
	for _, entity := range chunk.BlockEntities {
		if entity.X == x && entity.Y == y && entity.Z == z {
			return &entity, true
		}
	}
	return nil, false
}

// This function does not update the block for the players, so it should not be used. All of the position values should be absolute (aka (chunkPos<<4)+pos
func (chunk *Chunk) SetBlockEntity(x, y, z int32, be BlockEntity) {
	//c.LastModified = time.Now().UnixMilli()
	var index int = -1
	be.X, be.Y, be.Z = x, y, z
	for i, entity := range chunk.BlockEntities {
		if entity.X == x && entity.Y == y && entity.Z == z {
			index = i
			break
		}
	}
	if index == -1 {
		chunk.BlockEntities = append(chunk.BlockEntities, be)
		return
	}
	chunk.BlockEntities[index] = be
}
