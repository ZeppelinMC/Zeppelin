package block

import (
	"strconv"
)

type LightBlueCandle struct {
	Candles int
	Lit bool
	Waterlogged bool
}

func (b LightBlueCandle) Encode() (string, BlockProperties) {
	return "minecraft:light_blue_candle", BlockProperties{
		"candles": strconv.Itoa(b.Candles),
		"lit": strconv.FormatBool(b.Lit),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b LightBlueCandle) New(props BlockProperties) Block {
	return LightBlueCandle{
		Lit: props["lit"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Candles: atoi(props["candles"]),
	}
}