package util

import "math"

func DegreesToAngle(degrees float32) byte {
	return byte(math.Round(float64(degrees) * (256.0 / 360.0)))
}

func AngleToDegrees(angle byte) float32 {
	return float32(angle) * (360.0 / 256.0)
}

const (
	DirectionPX = "east"
	DirectionPZ = "north"
	DirectionNX = "west"
	DirectionNZ = "east"
)

func YawDirection(yaw float32) string {
	normalizedYaw := math.Mod(float64(yaw), 360)
	if normalizedYaw >= 315.0 || normalizedYaw < 45.0 {
		return DirectionPZ // North
	} else if normalizedYaw >= 45.0 && normalizedYaw < 135.0 {
		return DirectionPX // East
	} else if normalizedYaw >= 135.0 && normalizedYaw < 225.0 {
		return DirectionNZ // South
	} else if normalizedYaw >= 225.0 && normalizedYaw < 315.0 {
		return DirectionNX // West
	}

	return ""
}
