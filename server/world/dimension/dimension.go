package dimension

import (
	"fmt"
	"os"
	"sync"

	"github.com/zeppelinmc/zeppelin/server/session"
	"github.com/zeppelinmc/zeppelin/server/world/block/pos"
	"github.com/zeppelinmc/zeppelin/server/world/chunk"
	"github.com/zeppelinmc/zeppelin/server/world/chunk/section"
	"github.com/zeppelinmc/zeppelin/server/world/dimension/window"
	"github.com/zeppelinmc/zeppelin/server/world/region"
)

// if generateEmptyChunks is true, the server will override empty chunks with generated ones
func NewDimension(regionPath string, typ string, name string, broadcast *session.Broadcast, generator region.Generator, generateEmptyChunks bool) *Dimension {
	return &Dimension{
		regions: make(map[uint64]*region.File),

		regionPath:          regionPath,
		typ:                 typ,
		name:                name,
		generator:           generator,
		broadcast:           broadcast,
		WindowManager:       window.NewManager(),
		generateEmptyChunks: generateEmptyChunks && generator != nil,
	}
}

type Dimension struct {
	reg_mu  sync.Mutex
	regions map[uint64]*region.File

	broadcast *session.Broadcast

	generator           region.Generator
	generateEmptyChunks bool
	WindowManager       *window.WindowManager

	typ  string
	name string

	regionPath string
}

func (s *Dimension) SetBroadcast(b *session.Broadcast) {
	s.broadcast = b
}

func (s *Dimension) Type() string {
	return s.typ
}

func (s *Dimension) Name() string {
	return s.name
}

func (s *Dimension) SaveAllRegions() {
	s.reg_mu.Lock()
	defer s.reg_mu.Unlock()
	for hash, reg := range s.regions {
		rx, rz := s.regionUnhash(hash)
		path := fmt.Sprintf("%s/r.%d.%d.mca", s.regionPath, rx, rz)
		file, err := os.Create(path)
		if err != nil {
			continue
		}
		if reg.Encode(file) != nil {
			continue
		}
		if file.Close() != nil {
			continue
		}
	}
}

func (s *Dimension) LoadedChunks() int32 {
	s.reg_mu.Lock()
	defer s.reg_mu.Unlock()
	var count int32

	for _, reg := range s.regions {
		count += reg.LoadedChunks()
	}

	return count
}

func (s *Dimension) Block(x, y, z int32) (section.Block, error) {
	chunkX, chunkZ := x>>4, z>>4
	chunk, err := s.GetChunk(chunkX, chunkZ)
	if err != nil {
		return nil, err
	}
	return chunk.Block(x&0x0f, y, z&0x0f)
}

func (s *Dimension) SetBlock(pos pos.BlockPosition, b section.Block, placeSound bool) (state int64, err error) {
	chunkX, chunkZ := pos.ChunkX(), pos.ChunkZ()
	chunk, err := s.GetChunk(chunkX, chunkZ)
	if err != nil {
		return 0, err
	}
	i, err := chunk.SetBlock(pos.SectionX(), pos.Y(), pos.SectionZ(), b)
	if err != nil {
		return i, err
	}
	s.broadcast.UpdateBlock(pos, b, s.name, placeSound)
	return i, err
}

func (s *Dimension) SetBlockEntity(pos pos.BlockPosition, be chunk.BlockEntity) error {
	chunkX, chunkZ := pos.ChunkX(), pos.ChunkZ()
	chunk, err := s.GetChunk(chunkX, chunkZ)
	if err != nil {
		return err
	}
	chunk.SetBlockEntity(pos.X(), pos.Y(), pos.Z(), be)
	s.broadcast.UpdateBlockEntity(pos, be, s.name)

	return nil
}
func (s *Dimension) BlockEntity(x, y, z int32) (*chunk.BlockEntity, bool) {
	chunkX, chunkZ := x>>4, z>>4
	chunk, err := s.GetChunk(chunkX, chunkZ)
	if err != nil {
		return nil, false
	}
	return chunk.BlockEntity(x, y, z)
}

func (s *Dimension) GetChunk(x, z int32) (*chunk.Chunk, error) {
	rx, rz := s.chunkPosToRegionPos(x, z)
	region, err := s.getRegion(rx, rz)
	if err != nil {
		if s.generator != nil {
			region = s.newRegion(rx, rz)
		} else {
			return nil, err
		}
	}

	return region.GetChunk(x, z, s.generateEmptyChunks, s.generator)
}

func (s *Dimension) newRegion(rx, rz int32) *region.File {
	s.reg_mu.Lock()
	defer s.reg_mu.Unlock()
	hash := s.regionHash(rx, rz)
	s.regions[hash] = new(region.File)
	region.Empty(s.regions[hash])

	return s.regions[hash]
}

func (s *Dimension) getRegion(rx, rz int32) (*region.File, error) {
	s.reg_mu.Lock()
	defer s.reg_mu.Unlock()
	if r, ok := s.regions[s.regionHash(rx, rz)]; ok {
		return r, nil
	}

	reg, err := s.openRegion(rx, rz)
	if err != nil {
		return nil, err
	}

	return reg, err
}

func (s *Dimension) regionHash(rx, rz int32) uint64 {
	return uint64(uint32(rx)) | uint64(uint32(rz))<<32
}

func (s *Dimension) regionUnhash(hash uint64) (rx, rz int32) {
	urx := uint32(hash)
	urz := uint32(hash >> 32)

	return int32(urx), int32(urz)
}

func (s *Dimension) chunkPosToRegionPos(x, z int32) (rx, rz int32) {
	return x >> 5, z >> 5
}

func (s *Dimension) openRegion(rx, rz int32) (*region.File, error) {
	path := fmt.Sprintf("%s/r.%d.%d.mca", s.regionPath, rx, rz)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	//defer file.Close()
	hash := s.regionHash(rx, rz)

	s.regions[hash] = new(region.File)

	err = region.Decode(file, s.regions[hash])

	return s.regions[hash], err
}
