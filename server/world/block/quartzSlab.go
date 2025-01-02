package block

import (
	"strconv"
)

type QuartzSlab struct {
	Type string
	Waterlogged bool
}

func (b QuartzSlab) Encode() (string, BlockProperties) {
	return "minecraft:quartz_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b QuartzSlab) New(props BlockProperties) Block {
	return QuartzSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}