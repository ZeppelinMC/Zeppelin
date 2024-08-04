package block

import (
	"strconv"
)

type MediumAmethystBud struct {
	Facing string
	Waterlogged bool
}

func (b MediumAmethystBud) Encode() (string, BlockProperties) {
	return "minecraft:medium_amethyst_bud", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b MediumAmethystBud) New(props BlockProperties) Block {
	return MediumAmethystBud{
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}