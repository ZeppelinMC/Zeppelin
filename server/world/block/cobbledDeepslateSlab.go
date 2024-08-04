package block

import (
	"strconv"
)

type CobbledDeepslateSlab struct {
	Type string
	Waterlogged bool
}

func (b CobbledDeepslateSlab) Encode() (string, BlockProperties) {
	return "minecraft:cobbled_deepslate_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b CobbledDeepslateSlab) New(props BlockProperties) Block {
	return CobbledDeepslateSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}