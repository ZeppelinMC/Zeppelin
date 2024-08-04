package block

import (
	"strconv"
)

type LimeCandle struct {
	Candles int
	Lit bool
	Waterlogged bool
}

func (b LimeCandle) Encode() (string, BlockProperties) {
	return "minecraft:lime_candle", BlockProperties{
		"candles": strconv.Itoa(b.Candles),
		"lit": strconv.FormatBool(b.Lit),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b LimeCandle) New(props BlockProperties) Block {
	return LimeCandle{
		Candles: atoi(props["candles"]),
		Lit: props["lit"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}