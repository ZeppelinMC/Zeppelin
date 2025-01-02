package block

import (
	"strconv"
)

type PurpleCandleCake struct {
	Lit bool
}

func (b PurpleCandleCake) Encode() (string, BlockProperties) {
	return "minecraft:purple_candle_cake", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b PurpleCandleCake) New(props BlockProperties) Block {
	return PurpleCandleCake{
		Lit: props["lit"] != "false",
	}
}