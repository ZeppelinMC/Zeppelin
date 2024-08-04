package block

import (
	"strconv"
)

type GrayCandle struct {
	Lit bool
	Waterlogged bool
	Candles int
}

func (b GrayCandle) Encode() (string, BlockProperties) {
	return "minecraft:gray_candle", BlockProperties{
		"candles": strconv.Itoa(b.Candles),
		"lit": strconv.FormatBool(b.Lit),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b GrayCandle) New(props BlockProperties) Block {
	return GrayCandle{
		Candles: atoi(props["candles"]),
		Lit: props["lit"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}