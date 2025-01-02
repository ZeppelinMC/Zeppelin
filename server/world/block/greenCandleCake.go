package block

import (
	"strconv"
)

type GreenCandleCake struct {
	Lit bool
}

func (b GreenCandleCake) Encode() (string, BlockProperties) {
	return "minecraft:green_candle_cake", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b GreenCandleCake) New(props BlockProperties) Block {
	return GreenCandleCake{
		Lit: props["lit"] != "false",
	}
}