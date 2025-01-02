package dimension

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
	"github.com/zeppelinmc/zeppelin/server/world/chunk/section"
	"github.com/zeppelinmc/zeppelin/server/world/dimension/window"
	"github.com/zeppelinmc/zeppelin/server/world/level"
	"github.com/zeppelinmc/zeppelin/server/world/level/region"
	"github.com/zeppelinmc/zeppelin/util/log"
)

func New(regionPath, typ, name string, chunkGenerator chunk.Generator, level level.Level) *Dimension {
	return &Dimension{
		regions: make(map[uint64]*region.File),

		regionPath:    regionPath,
		typ:           typ,
		name:          name,
		generator:     chunkGenerator,
		WindowManager: window.NewManager(),
		Level:         level,
	}
}

type Dimension struct {
	reg_mu  sync.Mutex
	regions map[uint64]*region.File
	Level   level.Level

	generator     chunk.Generator
	WindowManager *window.WindowManager

	typ  string
	name string

	regionPath string
}

func (d *Dimension) Type() string {
	return d.typ
}

func (d *Dimension) Name() string {
	return d.name
}

func (d *Dimension) Save() {
	d.syncWindows()
	//s.saveAllRegions()
	d.Level.Close()
	log.Infoln("Saved dimension", d.name)
}

func (d *Dimension) syncWindows() {
	d.WindowManager.Range(func(i pos.BlockPosition, w *window.Window) {
		if w.ChunkEntityType != "" {
			d.SetBlockEntity(i, chunk.BlockEntity{
				X: i.X(), Y: i.Y(), Z: i.Z(),
				Id:    w.ChunkEntityType,
				Items: w.Items,
			})
		}
	})
}

func (d *Dimension) saveAllRegions() {
	d.reg_mu.Lock()
	defer d.reg_mu.Unlock()
	_ = os.MkdirAll(d.regionPath, 0755)
	for hash, reg := range d.regions {
		rx, rz := d.regionUnhash(hash)
		path := fmt.Sprintf("%s/r.%d.%d.mca", d.regionPath, rx, rz)
		file, err := os.Create(path)
		if err != nil {
			continue
		}
		if reg.Encode(file, region.CompressionZlib) != nil {
			continue
		}
		if file.Close() != nil {
			continue
		}
	}
}

func (d *Dimension) LoadedChunks() int32 {
	d.reg_mu.Lock()
	defer d.reg_mu.Unlock()
	var count int32

	for _, reg := range d.regions {
		count += reg.LoadedChunks()
	}

	return count
}

func (d *Dimension) Block(x, y, z int32) (section.AnvilBlock, error) {
	chunkX, chunkZ := x>>4, z>>4
	chunk, err := d.GetChunk(chunkX, chunkZ)
	if err != nil {
		return section.AnvilBlock{}, err
	}
	return chunk.Block(x&0x0f, y, z&0x0f)
}

func (d *Dimension) SetBlock(pos pos.BlockPosition, b section.AnvilBlock, placeSound bool) (state int64, err error) {
	chunkX, chunkZ := pos.ChunkX(), pos.ChunkZ()
	chunk, err := d.GetChunk(chunkX, chunkZ)
	if err != nil {
		return 0, err
	}
	i, err := chunk.SetBlock(pos.SectionX(), pos.Y(), pos.SectionZ(), b)
	if err != nil {
		return i, err
	}
	//s.broadcast.UpdateBlock(pos, b, s.name, placeSound)
	return i, err
}

func (d *Dimension) SetBlockEntity(pos pos.BlockPosition, be chunk.BlockEntity) error {
	chunkX, chunkZ := pos.ChunkX(), pos.ChunkZ()
	chunk, err := d.GetChunk(chunkX, chunkZ)
	if err != nil {
		return err
	}
	chunk.SetBlockEntity(pos.X(), pos.Y(), pos.Z(), be)
	//s.broadcast.UpdateBlockEntity(pos, be, s.name)

	return nil
}
func (d *Dimension) BlockEntity(x, y, z int32) (*chunk.BlockEntity, bool) {
	chunkX, chunkZ := x>>4, z>>4
	chunk, err := d.GetChunk(chunkX, chunkZ)
	if err != nil {
		return nil, false
	}
	return chunk.BlockEntity(x, y, z)
}

func (d *Dimension) GetChunkBuf(x, z int32, buf *bytes.Buffer) (*chunk.Chunk, error) {
	rx, rz := d.chunkPosToRegionPos(x, z)
	region, err := d.getRegion(rx, rz)
	if err != nil {
		if d.generator != nil {
			region = d.newRegion(rx, rz)
		} else {
			return nil, err
		}
	}

	return region.GetChunkBuf(x, z, buf)
}

func (d *Dimension) GetChunk(x, z int32) (*chunk.Chunk, error) {
	rx, rz := d.chunkPosToRegionPos(x, z)
	region, err := d.getRegion(rx, rz)
	if err != nil {
		if d.generator != nil {
			region = d.newRegion(rx, rz)
		} else {
			return nil, err
		}
	}

	return region.GetChunk(x, z)
}

func (d *Dimension) newRegion(rx, rz int32) *region.File {
	d.reg_mu.Lock()
	defer d.reg_mu.Unlock()
	hash := d.regionHash(rx, rz)
	d.regions[hash] = new(region.File)
	region.Empty(d.regions[hash], rx, rz, d.generator)

	return d.regions[hash]
}

func (d *Dimension) getRegion(rx, rz int32) (*region.File, error) {
	d.reg_mu.Lock()
	defer d.reg_mu.Unlock()
	if r, ok := d.regions[d.regionHash(rx, rz)]; ok {
		return r, nil
	}

	reg, err := d.openRegion(rx, rz)
	if err != nil {
		return nil, err
	}

	return reg, err
}

func (d *Dimension) regionHash(rx, rz int32) uint64 {
	return uint64(uint32(rx)) | uint64(uint32(rz))<<32
}

func (d *Dimension) regionUnhash(hash uint64) (rx, rz int32) {
	urx := uint32(hash)
	urz := uint32(hash >> 32)

	return int32(urx), int32(urz)
}

func (d *Dimension) chunkPosToRegionPos(x, z int32) (rx, rz int32) {
	return x >> 5, z >> 5
}

func (d *Dimension) openRegion(rx, rz int32) (*region.File, error) {
	path := fmt.Sprintf("%s/r.%d.%d.mca", d.regionPath, rx, rz)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	//defer file.Close()
	hash := d.regionHash(rx, rz)

	d.regions[hash] = new(region.File)

	err = region.Decode(file, d.regions[hash], rx, rz, d.generator)

	return d.regions[hash], err
}
