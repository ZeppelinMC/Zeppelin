package block

import (
	"strconv"
)

type SmallAmethystBud struct {
	Waterlogged bool
	Facing string
}

func (b SmallAmethystBud) Encode() (string, BlockProperties) {
	return "minecraft:small_amethyst_bud", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b SmallAmethystBud) New(props BlockProperties) Block {
	return SmallAmethystBud{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}