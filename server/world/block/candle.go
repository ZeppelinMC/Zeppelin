package block

import (
	"strconv"
)

type Candle struct {
	Lit bool
	Waterlogged bool
	Candles int
}

func (b Candle) Encode() (string, BlockProperties) {
	return "minecraft:candle", BlockProperties{
		"candles": strconv.Itoa(b.Candles),
		"lit": strconv.FormatBool(b.Lit),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b Candle) New(props BlockProperties) Block {
	return Candle{
		Candles: atoi(props["candles"]),
		Lit: props["lit"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}