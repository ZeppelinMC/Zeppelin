package block

import (
	"strconv"
)

type CobblestoneSlab struct {
	Waterlogged bool
	Type string
}

func (b CobblestoneSlab) Encode() (string, BlockProperties) {
	return "minecraft:cobblestone_slab", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"type": b.Type,
	}
}

func (b CobblestoneSlab) New(props BlockProperties) Block {
	return CobblestoneSlab{
		Waterlogged: props["waterlogged"] != "false",
		Type: props["type"],
	}
}