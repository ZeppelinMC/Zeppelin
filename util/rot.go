package util

import "math"

func DegreesToAngle(degrees float32) byte {
	return byte(math.Round(float64(degrees) * (256.0 / 360.0)))
}

func AngleToDegrees(angle byte) float32 {
	return float32(angle) * (360.0 / 256.0)
}
