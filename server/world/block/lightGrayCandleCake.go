package block

import (
	"strconv"
)

type LightGrayCandleCake struct {
	Lit bool
}

func (b LightGrayCandleCake) Encode() (string, BlockProperties) {
	return "minecraft:light_gray_candle_cake", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b LightGrayCandleCake) New(props BlockProperties) Block {
	return LightGrayCandleCake{
		Lit: props["lit"] != "false",
	}
}