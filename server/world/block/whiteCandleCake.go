package block

import (
	"strconv"
)

type WhiteCandleCake struct {
	Lit bool
}

func (b WhiteCandleCake) Encode() (string, BlockProperties) {
	return "minecraft:white_candle_cake", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b WhiteCandleCake) New(props BlockProperties) Block {
	return WhiteCandleCake{
		Lit: props["lit"] != "false",
	}
}