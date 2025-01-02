package block

import (
	"strconv"
)

type WhiteCandle struct {
	Candles int
	Lit bool
	Waterlogged bool
}

func (b WhiteCandle) Encode() (string, BlockProperties) {
	return "minecraft:white_candle", BlockProperties{
		"candles": strconv.Itoa(b.Candles),
		"lit": strconv.FormatBool(b.Lit),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b WhiteCandle) New(props BlockProperties) Block {
	return WhiteCandle{
		Lit: props["lit"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Candles: atoi(props["candles"]),
	}
}