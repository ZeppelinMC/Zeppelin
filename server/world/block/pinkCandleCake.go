package block

import (
	"strconv"
)

type PinkCandleCake struct {
	Lit bool
}

func (b PinkCandleCake) Encode() (string, BlockProperties) {
	return "minecraft:pink_candle_cake", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b PinkCandleCake) New(props BlockProperties) Block {
	return PinkCandleCake{
		Lit: props["lit"] != "false",
	}
}