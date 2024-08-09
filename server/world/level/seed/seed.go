// Package seed provides parsing of string seeds and generating them
package seed

import (
	"crypto/sha256"
	"encoding/binary"
	"math"
	"math/rand"
)

type Seed int64

// First 8 bytes of the SHA-256 hash of the world's seed. Used client side for biome noise
func (s Seed) Hash() int64 {
	hash := sha256.Sum256(binary.BigEndian.AppendUint64(nil, uint64(s)))

	return int64(binary.BigEndian.Uint64(hash[:8]))
}

func Random() Seed {
	return Seed(rand.Int63())
}

func New(str string) Seed {
	return Seed(hashCode(str))
}

// HashCode is an implementation of Java's hashCode function. It used to turn any string seed into a long seed
func hashCode(s string) int64 {
	var result int64
	n := len(s)

	for i := 0; i < len(s)-1; i++ {
		result += int64(s[i]) * int64(math.Pow(31, float64(n-(i+1))))
	}

	return result + int64(s[int(n)-1])
}
