package block

import (
	"strconv"
)

type MagentaCandleCake struct {
	Lit bool
}

func (b MagentaCandleCake) Encode() (string, BlockProperties) {
	return "minecraft:magenta_candle_cake", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b MagentaCandleCake) New(props BlockProperties) Block {
	return MagentaCandleCake{
		Lit: props["lit"] != "false",
	}
}