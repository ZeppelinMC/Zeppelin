package chunk

import "math"

type HeightMap struct {
	MotionBlocking []int64 `nbt:"MOTION_BLOCKING"`
	WorldSurface   []int64 `nbt:"WORLD_SURFACE"`
}

func (hm *HeightMap) SetMotionBlocking(x, z, y int) {
	hm.setState(hm.MotionBlocking, x, z, y)
}

func (hm *HeightMap) SetWorldSurface(x, z, y int) {
	hm.setState(hm.WorldSurface, x, z, y)
}

func (hm *HeightMap) GetWorldSurface(x, z int) int {
	return hm.getState(hm.WorldSurface, x, z)
}

func (hm *HeightMap) indexOffset(x, z int) (int, int) {
	blockNumber := z + x*16
	startLong := (blockNumber * 9) / 63
	stateOffset := (blockNumber * 9) % 63

	return startLong, stateOffset
}

func (hm *HeightMap) getState(s []int64, x, z int) int {
	i, offset := hm.indexOffset(x, z)

	states := s[i]
	states >>= offset

	data := states & (1<<9 - 1)

	return int(data)
}

func (hm *HeightMap) setState(s []int64, x, z, y int) {
	y += int(math.Abs(LowestY)) + 1

	i, offset := hm.indexOffset(x, z)
	mask := int64(^((1<<9 - 1) << offset))
	s[i] &= mask
	s[i] |= int64(y) << offset
}
