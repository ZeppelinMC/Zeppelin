package block

import (
	"strconv"
)

type GrayCandleCake struct {
	Lit bool
}

func (b GrayCandleCake) Encode() (string, BlockProperties) {
	return "minecraft:gray_candle_cake", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b GrayCandleCake) New(props BlockProperties) Block {
	return GrayCandleCake{
		Lit: props["lit"] != "false",
	}
}