package block

import (
	"strconv"
)

type SmoothQuartzSlab struct {
	Type string
	Waterlogged bool
}

func (b SmoothQuartzSlab) Encode() (string, BlockProperties) {
	return "minecraft:smooth_quartz_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b SmoothQuartzSlab) New(props BlockProperties) Block {
	return SmoothQuartzSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}