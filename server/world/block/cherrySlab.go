package block

import (
	"strconv"
)

type CherrySlab struct {
	Type string
	Waterlogged bool
}

func (b CherrySlab) Encode() (string, BlockProperties) {
	return "minecraft:cherry_slab", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"type": b.Type,
	}
}

func (b CherrySlab) New(props BlockProperties) Block {
	return CherrySlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}