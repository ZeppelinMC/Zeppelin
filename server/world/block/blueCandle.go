package block

import (
	"strconv"
)

type BlueCandle struct {
	Waterlogged bool
	Candles int
	Lit bool
}

func (b BlueCandle) Encode() (string, BlockProperties) {
	return "minecraft:blue_candle", BlockProperties{
		"lit": strconv.FormatBool(b.Lit),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"candles": strconv.Itoa(b.Candles),
	}
}

func (b BlueCandle) New(props BlockProperties) Block {
	return BlueCandle{
		Waterlogged: props["waterlogged"] != "false",
		Candles: atoi(props["candles"]),
		Lit: props["lit"] != "false",
	}
}