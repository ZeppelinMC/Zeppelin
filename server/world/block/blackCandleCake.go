package block

import (
	"strconv"
)

type BlackCandleCake struct {
	Lit bool
}

func (b BlackCandleCake) Encode() (string, BlockProperties) {
	return "minecraft:black_candle_cake", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b BlackCandleCake) New(props BlockProperties) Block {
	return BlackCandleCake{
		Lit: props["lit"] != "false",
	}
}