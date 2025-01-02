package block

import (
	"strconv"
)

type LightGrayBanner struct {
	Rotation int
}

func (b LightGrayBanner) Encode() (string, BlockProperties) {
	return "minecraft:light_gray_banner", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b LightGrayBanner) New(props BlockProperties) Block {
	return LightGrayBanner{
		Rotation: atoi(props["rotation"]),
	}
}