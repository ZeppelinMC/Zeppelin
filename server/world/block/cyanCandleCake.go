package block

import (
	"strconv"
)

type CyanCandleCake struct {
	Lit bool
}

func (b CyanCandleCake) Encode() (string, BlockProperties) {
	return "minecraft:cyan_candle_cake", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b CyanCandleCake) New(props BlockProperties) Block {
	return CyanCandleCake{
		Lit: props["lit"] != "false",
	}
}