package block

import (
	"strconv"
)

type RedSandstoneSlab struct {
	Type string
	Waterlogged bool
}

func (b RedSandstoneSlab) Encode() (string, BlockProperties) {
	return "minecraft:red_sandstone_slab", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"type": b.Type,
	}
}

func (b RedSandstoneSlab) New(props BlockProperties) Block {
	return RedSandstoneSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}