package block

import (
	"strconv"
)

type BrownCandle struct {
	Candles int
	Lit bool
	Waterlogged bool
}

func (b BrownCandle) Encode() (string, BlockProperties) {
	return "minecraft:brown_candle", BlockProperties{
		"candles": strconv.Itoa(b.Candles),
		"lit": strconv.FormatBool(b.Lit),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BrownCandle) New(props BlockProperties) Block {
	return BrownCandle{
		Candles: atoi(props["candles"]),
		Lit: props["lit"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}