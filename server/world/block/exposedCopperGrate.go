package block

import (
	"strconv"
)

type ExposedCopperGrate struct {
	Waterlogged bool
}

func (b ExposedCopperGrate) Encode() (string, BlockProperties) {
	return "minecraft:exposed_copper_grate", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b ExposedCopperGrate) New(props BlockProperties) Block {
	return ExposedCopperGrate{
		Waterlogged: props["waterlogged"] != "false",
	}
}