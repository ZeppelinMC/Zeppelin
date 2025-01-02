package block

import (
	"strconv"
)

type PetrifiedOakSlab struct {
	Type string
	Waterlogged bool
}

func (b PetrifiedOakSlab) Encode() (string, BlockProperties) {
	return "minecraft:petrified_oak_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b PetrifiedOakSlab) New(props BlockProperties) Block {
	return PetrifiedOakSlab{
		Waterlogged: props["waterlogged"] != "false",
		Type: props["type"],
	}
}