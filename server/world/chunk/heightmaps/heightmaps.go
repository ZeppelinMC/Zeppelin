package heightmaps

/* credit to https://github.com/aimjel for these calculations */

const MinY = -64

const HeightMapBitsPerEntry = 9

type Heightmap [37]int64

type Heightmaps struct {
	MotionBlocking Heightmap `nbt:"MOTION_BLOCKING"`
	WorldSurface   Heightmap `nbt:"WORLD_SURFACE"`
}

func (hm Heightmap) offset(x, z int32) (int32, int32) {
	blockNumber := z + x*16
	startLong := (blockNumber * HeightMapBitsPerEntry) / 63
	stateOffset := (blockNumber * HeightMapBitsPerEntry) % 63

	return startLong, stateOffset
}

func (hm Heightmap) Get(x, z int32) int32 {
	i, off := hm.offset(x, z)

	states := hm[i] >> off

	return int32(states & (1<<HeightMapBitsPerEntry - 1))
}

func (hm *Heightmap) Set(x, z, y int32) {
	y += -MinY + 1

	i, off := hm.offset(x, z)
	mask := int64(^((1<<HeightMapBitsPerEntry - 1) << off))
	//if len(*hm) <= int(i) {
	//	*hm = append(*hm, make(Heightmap, 37-len(*hm))...)
	//}
	(*hm)[i] &= mask
	(*hm)[i] |= int64(y) << off
}
