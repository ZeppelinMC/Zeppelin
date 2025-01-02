package block

import (
	"strconv"
)

type CrimsonSlab struct {
	Type string
	Waterlogged bool
}

func (b CrimsonSlab) Encode() (string, BlockProperties) {
	return "minecraft:crimson_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b CrimsonSlab) New(props BlockProperties) Block {
	return CrimsonSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}