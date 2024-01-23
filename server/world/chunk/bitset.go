package chunk

type bitSet struct {
	//the position of the bit we are writing at
	at int64

	out []int64

	//the index of the slice entry we edit
	i int
}

// set enables the xth bit in the set.
func (b *bitSet) set(x int) {
	if b.at == 64 {
		b.out = append(b.out, 0)
		b.i++
	}

	b.out[b.i] |= 1 << (x % 64)
	b.at++
}

func allZero(s []int8) bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}
	return true
}
