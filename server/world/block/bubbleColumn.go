package block

import (
	"strconv"
)

type BubbleColumn struct {
	Drag bool
}

func (b BubbleColumn) Encode() (string, BlockProperties) {
	return "minecraft:bubble_column", BlockProperties{
		"drag": strconv.FormatBool(b.Drag),
	}
}

func (b BubbleColumn) New(props BlockProperties) Block {
	return BubbleColumn{
		Drag: props["drag"] != "false",
	}
}