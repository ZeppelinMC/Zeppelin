package block

import (
	"strconv"
)

type BrownCandleCake struct {
	Lit bool
}

func (b BrownCandleCake) Encode() (string, BlockProperties) {
	return "minecraft:brown_candle_cake", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b BrownCandleCake) New(props BlockProperties) Block {
	return BrownCandleCake{
		Lit: props["lit"] != "false",
	}
}