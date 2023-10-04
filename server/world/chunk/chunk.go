package chunk

import (
	"errors"

	"github.com/aimjel/minecraft/nbt"
	"github.com/aimjel/minecraft/packet"
)

var ErrNotFound = errors.New("chunk not found")

var ErrIncomplete = errors.New("incomplete chunk found")

const lowestY = -64

type Chunk struct {
	x, z int32

	heightMap *HeightMap

	sections []*section
}

func NewAnvilChunk(b []byte) (*Chunk, error) {
	var ac anvilChunk
	if err := nbt.Unmarshal(b, &ac); err != nil {
		return nil, err
	}

	if ac.Status != "minecraft:full" {
		//TODO: create a chunk generator
		return nil, ErrIncomplete
	}

	var c = new(Chunk)
	c.x = ac.XPos
	c.z = ac.ZPos
	c.heightMap = &HeightMap{
		HeightMaps: struct {
			MotionBlocking []int64 `nbt:"MOTION_BLOCKING"`
			WorldSurface   []int64 `nbt:"WORLD_SURFACE"`
		}{ac.Heightmaps.MotionBlocking, ac.Heightmaps.WorldSurface}}

	c.sections = make([]*section, 0, len(ac.Sections))
	for _, s := range ac.Sections {
		if s.Y < 0 && s.Y < int8(ac.YPos) {
			continue
		}

		sec := newSection(s.BlockStates.Data, s.BlockStates.Palette, s.BlockLight, s.SkyLight)

		c.sections = append(c.sections, sec)
	}
	return c, nil
}

func (c *Chunk) Data() *packet.ChunkData {
	var pk packet.ChunkData
	pk.X, pk.Z = c.x, c.z
	pk.Heightmaps = (*c.heightMap).HeightMaps

	pk.Sections = make([]packet.Section, 0, len(c.sections)+2)
	for _, s := range c.sections {
		var sec packet.Section

		sec.BlockStates.Entries = s.ids
		sec.BlockStates.Data = s.data
		sec.BlockStates.BitsPerEntry = uint8(s.bitsPerEntry)
		sec.SkyLight = s.skyLight
		sec.BlockLight = s.blockLight
		pk.Sections = append(pk.Sections, sec)
	}

	return &pk
}

func HashXZ(x, z int32) uint64 {
	return (uint64(x) << 32) | uint64(z)
}
