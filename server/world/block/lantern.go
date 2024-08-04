package block

import (
	"strconv"
)

type Lantern struct {
	Hanging bool
	Waterlogged bool
}

func (b Lantern) Encode() (string, BlockProperties) {
	return "minecraft:lantern", BlockProperties{
		"hanging": strconv.FormatBool(b.Hanging),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b Lantern) New(props BlockProperties) Block {
	return Lantern{
		Hanging: props["hanging"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}