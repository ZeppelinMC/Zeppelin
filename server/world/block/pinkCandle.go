package block

import (
	"strconv"
)

type PinkCandle struct {
	Candles int
	Lit bool
	Waterlogged bool
}

func (b PinkCandle) Encode() (string, BlockProperties) {
	return "minecraft:pink_candle", BlockProperties{
		"candles": strconv.Itoa(b.Candles),
		"lit": strconv.FormatBool(b.Lit),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b PinkCandle) New(props BlockProperties) Block {
	return PinkCandle{
		Lit: props["lit"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Candles: atoi(props["candles"]),
	}
}