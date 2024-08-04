package block

import (
	"strconv"
)

type CandleCake struct {
	Lit bool
}

func (b CandleCake) Encode() (string, BlockProperties) {
	return "minecraft:candle_cake", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b CandleCake) New(props BlockProperties) Block {
	return CandleCake{
		Lit: props["lit"] != "false",
	}
}