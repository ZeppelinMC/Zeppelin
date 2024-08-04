package block

import (
	"strconv"
)

type RedCandle struct {
	Lit bool
	Waterlogged bool
	Candles int
}

func (b RedCandle) Encode() (string, BlockProperties) {
	return "minecraft:red_candle", BlockProperties{
		"candles": strconv.Itoa(b.Candles),
		"lit": strconv.FormatBool(b.Lit),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b RedCandle) New(props BlockProperties) Block {
	return RedCandle{
		Lit: props["lit"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Candles: atoi(props["candles"]),
	}
}