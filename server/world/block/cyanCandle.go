package block

import (
	"strconv"
)

type CyanCandle struct {
	Candles int
	Lit bool
	Waterlogged bool
}

func (b CyanCandle) Encode() (string, BlockProperties) {
	return "minecraft:cyan_candle", BlockProperties{
		"candles": strconv.Itoa(b.Candles),
		"lit": strconv.FormatBool(b.Lit),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b CyanCandle) New(props BlockProperties) Block {
	return CyanCandle{
		Waterlogged: props["waterlogged"] != "false",
		Candles: atoi(props["candles"]),
		Lit: props["lit"] != "false",
	}
}