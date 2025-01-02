package block

import (
	"strconv"
)

type LimeCandleCake struct {
	Lit bool
}

func (b LimeCandleCake) Encode() (string, BlockProperties) {
	return "minecraft:lime_candle_cake", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b LimeCandleCake) New(props BlockProperties) Block {
	return LimeCandleCake{
		Lit: props["lit"] != "false",
	}
}