package player

import "strings"

func Gamemode(gm string) int {
	switch strings.ToLower(gm) {
	case "survival":
		return 0
	case "creative":
		return 1
	case "adventure":
		return 2
	case "spectator":
		return 3
	}

	return -1
}

func GamemodeName(gm byte) string {
	switch gm {
	case 0:
		return "survival"
	case 1:
		return "creative"
	case 2:
		return "adventure"
	case 3:
		return "spectator"
	default:
		return ""
	}
}
