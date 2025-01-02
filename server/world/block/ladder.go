package block

import (
	"strconv"
)

type Ladder struct {
	Facing string
	Waterlogged bool
}

func (b Ladder) Encode() (string, BlockProperties) {
	return "minecraft:ladder", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b Ladder) New(props BlockProperties) Block {
	return Ladder{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}