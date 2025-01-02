package block

import (
	"strconv"
)

type BlueCandleCake struct {
	Lit bool
}

func (b BlueCandleCake) Encode() (string, BlockProperties) {
	return "minecraft:blue_candle_cake", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b BlueCandleCake) New(props BlockProperties) Block {
	return BlueCandleCake{
		Lit: props["lit"] != "false",
	}
}