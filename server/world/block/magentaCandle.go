package block

import (
	"strconv"
)

type MagentaCandle struct {
	Candles int
	Lit bool
	Waterlogged bool
}

func (b MagentaCandle) Encode() (string, BlockProperties) {
	return "minecraft:magenta_candle", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"candles": strconv.Itoa(b.Candles),
		"lit": strconv.FormatBool(b.Lit),
	}
}

func (b MagentaCandle) New(props BlockProperties) Block {
	return MagentaCandle{
		Lit: props["lit"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Candles: atoi(props["candles"]),
	}
}