package chunk

import (
	"bytes"
	"errors"
	"github.com/aimjel/minecraft/nbt"
	"github.com/aimjel/minecraft/protocol/encoding"
	"github.com/aimjel/nitrate/server/block"
	"math"
	"unsafe"
)

var ErrNotFound = errors.New("chunk not found")

const LowestY = -64

// yOffset uses the smallest y value to create an offset,
// so every Y Section can be indexed within a slice.
const yOffset = (LowestY / 16) * -1

// Chunk represents a full minecraft chunk. 16x384x16.
type Chunk struct {
	X, Z int32

	HeightMap *HeightMap

	sections []*Section
}

var ErrIncomplete = errors.New("chunk is incomplete")

func New(x, z int32) *Chunk {
	c := &Chunk{
		X: x,
		Z: z,
		HeightMap: &HeightMap{
			MotionBlocking: make([]int64, 37),
			WorldSurface:   make([]int64, 37),
		},
		sections: make([]*Section, 0, 24),
	}

	//you need to send all chunks section, even if they are empty.
	for i := 0; i < 24; i++ {
		c.sections = append(c.sections, NewSection([]block.Block{block.Air{}}, nil, nil, nil))
	}

	return c
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
		X:         ac.XPos,
		Z:         ac.ZPos,
		HeightMap: &ac.Heightmaps,
	}
	c.sections = make([]*Section, 0, len(ac.Sections))
	for _, s := range ac.Sections {

		blocks := make([]block.Block, 0, len(s.BlockStates.Palette))
		for _, entry := range s.BlockStates.Palette {
			bl := block.GetBlock(entry.Name)
			if entry.Properties != nil {
				bl = bl.New(entry.Properties)
			}
			blocks = append(blocks, bl)
		}

		c.sections = append(c.sections, NewSection(blocks, s.BlockStates.Data, s.SkyLight, s.BlockLight))
	}

	return c, nil
}

func (c *Chunk) SetBlock(x, y, z int, b block.Block) {
	x, z = 0xf&x, 0xf&z

	ySec := int(math.Floor(float64(y)/16) + yOffset)

	if len(c.sections) <= ySec {
		for len(c.sections) <= ySec {
			c.sections = append(c.sections, NewSection([]block.Block{block.Air{}}, nil, nil, nil))
		}
	}

	c.sections[ySec].SetBlock(x, y&0xf, z, b)
}

func (c *Chunk) NetworkEncode(buf *bytes.Buffer) {
	//heightmap
	//sections
	//todo block entities
	//lights

	pw := encoding.NewWriter(buf)

	err := nbt.NewEncoder(buf).Encode(*c.HeightMap)
	if err != nil {
		panic(err)
	}

	//5 is the max bytes for a var-int
	buf.Write([]byte{1, 2, 3, 4, 5})
	start := buf.Len()
	for _, sec := range c.sections {
		//fmt.Printf("%v %#v\n\n", i, sec)
		if len(sec.blocks) == 1 && sec.blocks[0].EncodedName() == "minecraft:air" {
			_ = pw.Uint16(0)
		} else {
			_ = pw.Uint16(1024) //block count todo implement properly
		}

		//encodes the palette container for block states
		_ = pw.Uint8(uint8(sec.bitsPerEntry))
		if sec.bitsPerEntry != 0 {
			//encodes indirect palettes
			_ = pw.VarInt(int32(len(sec.ids)))
		}

		for _, id := range sec.ids {
			_ = pw.VarInt(id)
		}
		_ = pw.Int64Array(sec.data)

		_ = pw.Uint8(0)    //bits per entry for biomes
		_ = pw.Uint8(0x39) //used one biome id, for single value palette, //todo implement proper biome id
		_ = pw.Uint8(0)    //length of the data array
	}

	secData := buf.Bytes()[start:buf.Len()]
	buf.Truncate(start - 5)

	_ = pw.VarInt(int32(len(secData)))
	buf.Write(secData)

	//todo block entities
	_ = pw.Uint8(0)

	c.NetworkEncodeLight(buf)
}

// lightBuf represents the light data of a chunk section with no skylight.
// Used for the section under the bedrock
var lightBuf = make([]byte, 2048)

func (c *Chunk) NetworkEncodeLight(buf *bytes.Buffer) {
	w := encoding.NewWriter(buf)

	skyLight := bitSet{out: make([]int64, 1)}
	skyLightEmpty := bitSet{out: make([]int64, 1)}
	skyArrays := uint8(1) //account for the light data under the bedrock

	//bit 0 is the section under the bedrock.
	//Yes, minecraft cares about the light data under the bedrock.
	skyLight.set(0)
	skyLightEmpty.set(0)

	for i, sec := range c.sections {
		if sec.skyLight == nil {
			continue
		}

		skyLight.set(i + 1)

		if allZero(sec.skyLight) {
			skyLightEmpty.set(i + 1)
		}

		skyArrays++
	}

	_ = w.Int64Array(skyLight.out)
	_ = w.Int64Array([]int64{0}) //todo block light mask

	_ = w.Int64Array(skyLightEmpty.out)
	_ = w.Int64Array([]int64{}) // todo empty block-light mask

	_ = w.Uint8(skyArrays)
	//this sends the skylight data under the bedrock
	_ = w.ByteArray(lightBuf)
	for _, sec := range c.sections {
		if sec.skyLight == nil {
			continue
		}

		//fmt.Printf("%#v\n", sec)
		_ = w.ByteArray(*(*[]uint8)(unsafe.Pointer(&sec.skyLight)))
	}

	_ = w.Uint8(0) //todo no block light info lol
}

func (c *Chunk) EncodeHeightMap(buf *bytes.Buffer) {
	if err := nbt.NewEncoder(buf).Encode(*c.HeightMap); err != nil {
		panic(err)
	}
}

// GenerateSkyLight generates skylight data for all sections
func (c *Chunk) GenerateSkyLight() {
	var prevEmpty bool
	//searches through the sections from top to bottom
	// finding an empty chunk first and then a non-empty chunk
	for i := len(c.sections) - 1; i >= 0; i-- {
		if len(c.sections[i].blocks) == 1 && c.sections[i].blocks[0].EncodedName() == "minecraft:air" {
			prevEmpty = true
			continue
		}

		if prevEmpty {
			//loop over every (x,z) coordinate
			//get the highest block using the heightmap
			//then set skylight
			for x := 0; x < 16; x++ {
				for z := 0; z < 16; z++ {
					// the highest point (x, y, z)
					y := c.HeightMap.GetWorldSurface(x, z)

					//ySec := ((y - 64) / 16) + yOffset

					//fmt.Println("HIGHEST BLOCK (y,x,z)", y+LowestY, x, z, "chunk section", ySec)

					//fmt.Println("j=", y-64, "; j<", ((i+1)-yOffset)*16)
					for j := y - 64; j < ((i+1)-yOffset)*16; j++ {
						c.SetSkyLight(x, j, z, 15)
					}
				}
			}

			//fmt.Println(i+1, "adding full skylight data for section")
			if c.sections[i+1].skyLight == nil {
				c.sections[i+1].skyLight = make([]int8, 2048)
			}

			for n := range c.sections[i+1].skyLight {
				c.sections[i+1].skyLight[n] = -1
			}

			return
		}
	}
}

func (c *Chunk) SetSkyLight(x, y, z int, level uint8) {
	ySec := int(math.Floor(float64(y)/16) + yOffset)
	//fmt.Println("editing skylight in section", ySec, (float64(y)/16)+yOffset)
	c.sections[ySec].setSkyLight(x, y&0xf, z, level)
}

func Hash(x, z int32) uint64 {
	return uint64(uint32(x))<<32 | uint64(uint32(z))
}
