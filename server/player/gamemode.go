package player

import (
	"strings"
)

type gameMode int32

const (
	Unknown gameMode = iota - 1
	Survival
	Creative
	Adventure
	Spectator
)

func (gm gameMode) String() string {
	switch gm {
	case Survival:
		return "survival"

	case Creative:
		return "creative"

	case Adventure:
		return "adventure"

	case Spectator:
		return "spectator"
	}

	return "unknown"
}

func GameMode[T int | int32](gm T) gameMode {
	switch gm {
	case 0, 1, 2, 3:
		return gameMode(gm)
	}

	return Unknown
}

// GameModeAtoi converts a game mode string into a game mode type
func GameModeAtoi(gm string) gameMode {
	switch strings.ToLower(gm) {

	case "survival", "s":
		return Survival

	case "creative", "c":
		return Creative

	case "adventure":
		return Adventure

	case "spectator":
		return Spectator
	}

	return Unknown
}
