package chunk

import (
	"errors"
	"fmt"

	"github.com/aimjel/minecraft/protocol/types"
	"github.com/dynamitemc/dynamite/logger"

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

func (c *Chunk) SetPosition(x, z int32) {
	c.x, c.z = x, z
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

	// if c.heightMap != nil {
	// 	pk.Heightmaps = *c.heightMap
	// }

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

		logger.Println("data length", len(sec.BlockStates.Data))
		logger.Println("entries length", len(sec.BlockStates.Entries))
	}

	return &pk
}

func (c *Chunk) Block(x, y, z int64) Block {
	relx, rely, relz := x&0x0f, y&0x0f, z&0x0f
	if y < lowestY {
		return GetBlock("minecraft:air")
	}

	sec := c.ySection(rely)
	b := sec.GetBlockAt(int(relx), int(rely), int(relz))
	//logger.Println(b.EncodedName())
	return b
}

func (c *Chunk) SetBlock(x, y, z int64, b Block) {
	relx, rely, relz := x&0x0f, y&0x0f, z&0x0f

	sec := c.ySection(rely)
	sec.setBlockAt(int(relx), int(rely), int(relz), b)
}

func (c *Chunk) ySection(y int64) *section {
	ySecIndex := int(y/16) + 4
	if ySecIndex >= len(c.sections) {
		fill := ySecIndex - len(c.sections)
		//adds all the slices inbetween
		for i := 0; i < fill+1; i++ {
			c.sections = append(c.sections, newSection(nil, []blockEntry{{Name: "minecraft:air"}}, nil, nil))
		}

		fmt.Println(fill+1, "sections created")
	}

	return c.sections[ySecIndex]
}

func (c *Chunk) RandomTick(speed int32) {
	for i := int32(0); i < speed; i++ {
		//c.Block()
	}
}

func (c *Chunk) Sections() []*section {
	return c.sections
}
