package block

import (
	"strconv"
)

type PurpleCandle struct {
	Candles int
	Lit bool
	Waterlogged bool
}

func (b PurpleCandle) Encode() (string, BlockProperties) {
	return "minecraft:purple_candle", BlockProperties{
		"candles": strconv.Itoa(b.Candles),
		"lit": strconv.FormatBool(b.Lit),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b PurpleCandle) New(props BlockProperties) Block {
	return PurpleCandle{
		Candles: atoi(props["candles"]),
		Lit: props["lit"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}