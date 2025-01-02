package block

import (
	"strconv"
)

type BigDripleafStem struct {
	Facing string
	Waterlogged bool
}

func (b BigDripleafStem) Encode() (string, BlockProperties) {
	return "minecraft:big_dripleaf_stem", BlockProperties{
		"facing": b.Facing,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BigDripleafStem) New(props BlockProperties) Block {
	return BigDripleafStem{
		Facing: props["facing"],
		Waterlogged: props["waterlogged"] != "false",
	}
}