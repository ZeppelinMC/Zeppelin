package block

import (
	"strconv"
)

type LargeAmethystBud struct {
	Facing string
	Waterlogged bool
}

func (b LargeAmethystBud) Encode() (string, BlockProperties) {
	return "minecraft:large_amethyst_bud", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b LargeAmethystBud) New(props BlockProperties) Block {
	return LargeAmethystBud{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}