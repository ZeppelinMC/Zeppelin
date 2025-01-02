package block

import (
	"strconv"
)

type BrainCoral struct {
	Waterlogged bool
}

func (b BrainCoral) Encode() (string, BlockProperties) {
	return "minecraft:brain_coral", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BrainCoral) New(props BlockProperties) Block {
	return BrainCoral{
		Waterlogged: props["waterlogged"] != "false",
	}
}