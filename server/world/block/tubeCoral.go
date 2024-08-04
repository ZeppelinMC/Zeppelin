package block

import (
	"strconv"
)

type TubeCoral struct {
	Waterlogged bool
}

func (b TubeCoral) Encode() (string, BlockProperties) {
	return "minecraft:tube_coral", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b TubeCoral) New(props BlockProperties) Block {
	return TubeCoral{
		Waterlogged: props["waterlogged"] != "false",
	}
}