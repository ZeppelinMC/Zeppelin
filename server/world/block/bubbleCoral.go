package block

import (
	"strconv"
)

type BubbleCoral struct {
	Waterlogged bool
}

func (b BubbleCoral) Encode() (string, BlockProperties) {
	return "minecraft:bubble_coral", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b BubbleCoral) New(props BlockProperties) Block {
	return BubbleCoral{
		Waterlogged: props["waterlogged"] != "false",
	}
}