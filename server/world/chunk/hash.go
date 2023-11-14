package chunk

type Hash uint64

func HashXZ(x, z int32) Hash {
	return Hash(uint32(x))<<32 | Hash(uint32(z))
}

func (h Hash) Position() (x, z int32) {
	z = int32(h & 0xFFFFFFFF)
	x = int32((h >> 32) & 0xFFFFFFFF)
	return x, z
}
