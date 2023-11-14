package chunk

import (
	"errors"

	"github.com/aimjel/minecraft/protocol/types"

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

	Sections []*section
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

	c.Sections = make([]*section, 0, len(ac.Sections))
	for _, s := range ac.Sections {
		if s.Y < 0 && s.Y < int8(ac.YPos) {
			continue
		}
		sec := newSection(s.BlockStates.Data, s.BlockStates.Palette, s.BlockLight, s.SkyLight)

		c.Sections = append(c.Sections, sec)
	}
	return c, nil
}

func (c *Chunk) Data() *packet.ChunkData {
	var pk packet.ChunkData
	pk.X, pk.Z = c.x, c.z
	pk.Heightmaps = *c.heightMap

	pk.Sections = make([]types.ChunkSection, 0, len(c.Sections)+2)
	for _, s := range c.Sections {
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

func (c *Chunk) Block(x, y, z int64) Block {
	relx, rely, relz := x&0x0f, y&0x0f, z&0x0f
	if y < lowestY {
		return GetBlock("minecraft:air")
	}

	sec := c.Sections[int(y/16)+4]
	b := sec.GetBlockAt(int(relx), int(rely), int(relz))
	//logger.Println(b.EncodedName())
	return b
}

func (c *Chunk) SetBlock(x, y, z int64, b Block) {
	y1 := int(y/16) + 4
	relx, rely, relz := x&0x0f, y&0x0f, z&0x0f

	sec := c.Sections[y1]
	sec.setBlockAt(int(relx), int(rely), int(relz), b)
}

func (c *Chunk) RandomTick(speed int32) {
	for i := int32(0); i < speed; i++ {
		//c.Block()
	}
}
