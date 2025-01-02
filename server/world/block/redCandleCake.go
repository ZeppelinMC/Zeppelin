package block

import (
	"strconv"
)

type RedCandleCake struct {
	Lit bool
}

func (b RedCandleCake) Encode() (string, BlockProperties) {
	return "minecraft:red_candle_cake", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b RedCandleCake) New(props BlockProperties) Block {
	return RedCandleCake{
		Lit: props["lit"] != "false",
	}
}