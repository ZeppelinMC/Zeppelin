package block

import (
	"strconv"
)

type YellowCandleCake struct {
	Lit bool
}

func (b YellowCandleCake) Encode() (string, BlockProperties) {
	return "minecraft:yellow_candle_cake", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b YellowCandleCake) New(props BlockProperties) Block {
	return YellowCandleCake{
		Lit: props["lit"] != "false",
	}
}