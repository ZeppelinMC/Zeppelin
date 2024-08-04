package block

import (
	"strconv"
)

type LightBlueCandleCake struct {
	Lit bool
}

func (b LightBlueCandleCake) Encode() (string, BlockProperties) {
	return "minecraft:light_blue_candle_cake", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b LightBlueCandleCake) New(props BlockProperties) Block {
	return LightBlueCandleCake{
		Lit: props["lit"] != "false",
	}
}