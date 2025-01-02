package block

import (
	"strconv"
)

type GreenCandle struct {
	Candles int
	Lit bool
	Waterlogged bool
}

func (b GreenCandle) Encode() (string, BlockProperties) {
	return "minecraft:green_candle", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"candles": strconv.Itoa(b.Candles),
	}
}

func (b GreenCandle) New(props BlockProperties) Block {
	return GreenCandle{
		Candles: atoi(props["candles"]),
		Lit: props["lit"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}