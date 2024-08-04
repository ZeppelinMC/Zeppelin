package block

import (
	"strconv"
)

type YellowCandle struct {
	Candles int
	Lit bool
	Waterlogged bool
}

func (b YellowCandle) Encode() (string, BlockProperties) {
	return "minecraft:yellow_candle", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"candles": strconv.Itoa(b.Candles),
	}
}

func (b YellowCandle) New(props BlockProperties) Block {
	return YellowCandle{
		Candles: atoi(props["candles"]),
		Lit: props["lit"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}