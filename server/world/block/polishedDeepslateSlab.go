package block

import (
	"strconv"
)

type PolishedDeepslateSlab struct {
	Waterlogged bool
	Type string
}

func (b PolishedDeepslateSlab) Encode() (string, BlockProperties) {
	return "minecraft:polished_deepslate_slab", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"type": b.Type,
	}
}

func (b PolishedDeepslateSlab) New(props BlockProperties) Block {
	return PolishedDeepslateSlab{
		Waterlogged: props["waterlogged"] != "false",
		Type: props["type"],
	}
}