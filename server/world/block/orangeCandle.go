package block

import (
	"strconv"
)

type OrangeCandle struct {
	Candles int
	Lit bool
	Waterlogged bool
}

func (b OrangeCandle) Encode() (string, BlockProperties) {
	return "minecraft:orange_candle", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"candles": strconv.Itoa(b.Candles),
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b OrangeCandle) New(props BlockProperties) Block {
	return OrangeCandle{
		Candles: atoi(props["candles"]),
		Lit: props["lit"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}