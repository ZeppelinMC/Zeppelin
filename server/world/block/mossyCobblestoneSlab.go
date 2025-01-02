package block

import (
	"strconv"
)

type MossyCobblestoneSlab struct {
	Waterlogged bool
	Type string
}

func (b MossyCobblestoneSlab) Encode() (string, BlockProperties) {
	return "minecraft:mossy_cobblestone_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b MossyCobblestoneSlab) New(props BlockProperties) Block {
	return MossyCobblestoneSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}