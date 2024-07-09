package region

import (
	"fmt"
	"math"
	"os"
	"sync"
)

func NewSave(regionPath string) *Save {
	return &Save{
		regions: make(map[uint64]*RegionFile),

		regionPath: regionPath,
	}
}

type Save struct {
	reg_mu  sync.Mutex
	regions map[uint64]*RegionFile

	regionPath string
}

func (s *Save) GetChunk(x, z int32) (*Chunk, error) {
	region, err := s.getRegion(s.chunkPosToRegionPos(x, z))
	if err != nil {
		return nil, err
	}

	return region.GetChunk(x, z)
}

func (s *Save) getRegion(rx, rz int32) (*RegionFile, error) {
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

func (s *Save) regionHash(rx, rz int32) uint64 {
	return uint64(rx) | uint64(rz)<<32
}

func (s *Save) chunkPosToRegionPos(x, z int32) (rx, rz int32) {
	return int32(math.Floor(float64(x) / 32)), int32(math.Floor(float64(z) / 32))
}

func (s *Save) openRegion(rx, rz int32) (*RegionFile, error) {
	path := fmt.Sprintf("%s/r.%d.%d.mca", s.regionPath, rx, rz)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	//defer file.Close()
	hash := s.regionHash(rx, rz)

	s.regions[hash] = new(RegionFile)

	err = DecodeRegion(file, s.regions[hash])

	return s.regions[hash], err
}
