package block

import (
	"strconv"
)

type SmallDripleaf struct {
	Facing string
	Half string
	Waterlogged bool
}

func (b SmallDripleaf) Encode() (string, BlockProperties) {
	return "minecraft:small_dripleaf", BlockProperties{
		"half": b.Half,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
	}
}

func (b SmallDripleaf) New(props BlockProperties) Block {
	return SmallDripleaf{
		Facing: props["facing"],
		Half: props["half"],
		Waterlogged: props["waterlogged"] != "false",
	}
}