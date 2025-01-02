// Package uuid provides a type for UUIDs stored in level files

package uuid

import "github.com/google/uuid"

func New(u uuid.UUID) UUID {
	return UUID{
		int32(u[0])<<24 | int32(u[1])<<16 | int32(u[2])<<8 | int32(u[3]),
		int32(u[4])<<24 | int32(u[5])<<16 | int32(u[6])<<8 | int32(u[7]),
		int32(u[8])<<24 | int32(u[9])<<16 | int32(u[10])<<8 | int32(u[11]),
		int32(u[12])<<24 | int32(u[13])<<16 | int32(u[14])<<8 | int32(u[15]),
	}
}

// A uuid saved in playerdata and state files contains 4 integers of the 128 bit uuid ordered from most to least significant
type UUID [4]int32

// Returns the UUID object of this data uuid
func (u UUID) UUID() uuid.UUID {
	return uuid.UUID{
		byte(u[0] >> 24),
		byte(u[0] >> 16),
		byte(u[0] >> 8),
		byte(u[0]),

		byte(u[1] >> 24),
		byte(u[1] >> 16),
		byte(u[1] >> 8),
		byte(u[1]),

		byte(u[2] >> 24),
		byte(u[2] >> 16),
		byte(u[2] >> 8),
		byte(u[2]),

		byte(u[3] >> 24),
		byte(u[3] >> 16),
		byte(u[3] >> 8),
		byte(u[3]),
	}
}
