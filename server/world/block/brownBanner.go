package block

import (
	"strconv"
)

type BrownBanner struct {
	Rotation int
}

func (b BrownBanner) Encode() (string, BlockProperties) {
	return "minecraft:brown_banner", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b BrownBanner) New(props BlockProperties) Block {
	return BrownBanner{
		Rotation: atoi(props["rotation"]),
	}
}