package block

import (
	"strconv"
)

type AmethystCluster struct {
	Waterlogged bool
	Facing string
}

func (b AmethystCluster) Encode() (string, BlockProperties) {
	return "minecraft:amethyst_cluster", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b AmethystCluster) New(props BlockProperties) Block {
	return AmethystCluster{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}