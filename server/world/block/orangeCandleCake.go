package block

import (
	"strconv"
)

type OrangeCandleCake struct {
	Lit bool
}

func (b OrangeCandleCake) Encode() (string, BlockProperties) {
	return "minecraft:orange_candle_cake", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b OrangeCandleCake) New(props BlockProperties) Block {
	return OrangeCandleCake{
		Lit: props["lit"] != "false",
	}
}