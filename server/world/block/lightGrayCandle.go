package block

import (
	"strconv"
)

type LightGrayCandle struct {
	Candles int
	Lit bool
	Waterlogged bool
}

func (b LightGrayCandle) Encode() (string, BlockProperties) {
	return "minecraft:light_gray_candle", BlockProperties{
		"candles": strconv.Itoa(b.Candles),
		"lit": strconv.FormatBool(b.Lit),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b LightGrayCandle) New(props BlockProperties) Block {
	return LightGrayCandle{
		Candles: atoi(props["candles"]),
		Lit: props["lit"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}