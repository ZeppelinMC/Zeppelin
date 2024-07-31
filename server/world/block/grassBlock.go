package block

import (
	"strconv"
)

type GrassBlock struct {
	Snowy bool
}

func (g GrassBlock) Encode() (string, BlockProperties) {
	return "minecraft:grass_block", BlockProperties{
		"snowy": strconv.FormatBool(g.Snowy),
	}
}

func (g GrassBlock) New(props BlockProperties) Block {
	return GrassBlock{Snowy: props["snowy"] == "true"}
}

var _ Block = (*GrassBlock)(nil)
