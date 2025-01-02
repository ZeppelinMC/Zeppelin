package block

import (
	"strconv"
)

type BlackCandle struct {
	Candles int
	Lit bool
	Waterlogged bool
}

func (b BlackCandle) Encode() (string, BlockProperties) {
	return "minecraft:black_candle", BlockProperties{
		"candles": strconv.Itoa(b.Candles),
		"lit": strconv.FormatBool(b.Lit),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BlackCandle) New(props BlockProperties) Block {
	return BlackCandle{
		Candles: atoi(props["candles"]),
		Lit: props["lit"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}