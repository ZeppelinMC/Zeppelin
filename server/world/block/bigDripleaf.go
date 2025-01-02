package block

import (
	"strconv"
)

type BigDripleaf struct {
	Tilt string
	Waterlogged bool
	Facing string
}

func (b BigDripleaf) Encode() (string, BlockProperties) {
	return "minecraft:big_dripleaf", BlockProperties{
		"facing": b.Facing,
		"tilt": b.Tilt,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BigDripleaf) New(props BlockProperties) Block {
	return BigDripleaf{
		Facing: props["facing"],
		Tilt: props["tilt"],
		Waterlogged: props["waterlogged"] != "false",
	}
}