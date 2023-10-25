package chunk

import (
	"errors"
	"fmt"

	"github.com/aimjel/minecraft/protocol/types"
	"github.com/dynamitemc/dynamite/server/block"

	"github.com/aimjel/minecraft/nbt"
	"github.com/aimjel/minecraft/packet"
)

var ErrNotFound = errors.New("chunk not found")

var ErrIncomplete = errors.New("incomplete chunk found")

const lowestY = -64

type Chunk struct {
	x, z int32

	heightMap *HeightMap

	Entities []Entity

	sections []*section
}

func NewAnvilChunk(b []byte) (*Chunk, error) {
	var ac anvilChunk
	if err := nbt.Unmarshal(b, &ac); err != nil {
		return nil, err
	}

	if ac.Status != "minecraft:full" {
		return nil, ErrIncomplete
	}

	c := &Chunk{
		x: ac.XPos,
		z: ac.ZPos,
		heightMap: &HeightMap{
			MotionBlocking: ac.Heightmaps.MotionBlocking,
			WorldSurface:   ac.Heightmaps.WorldSurface,
		},
	}

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
	pk.Heightmaps = *c.heightMap

	pk.Sections = make([]types.ChunkSection, 0, len(c.sections)+2)
	for _, s := range c.sections {
		if s == nil {
			continue
		}

		var sec types.ChunkSection

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
	return uint64(uint32(x))<<32 | uint64(uint32(z))
}

func (c *Chunk) Block(x, y, z int64) block.Block {
	y1 := int(y/16) + 4
	relx, rely, relz := x&0x0f, y&0x0f, z&0x0f

	sec := c.sections[y1]
	fmt.Println(sec.getBlockAt(int(relx), int(rely), int(relz)).EncodedName())
	return nil
}
