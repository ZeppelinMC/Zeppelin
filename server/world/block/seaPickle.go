package block

import (
	"strconv"
)

type SeaPickle struct {
	Pickles int
	Waterlogged bool
}

func (b SeaPickle) Encode() (string, BlockProperties) {
	return "minecraft:sea_pickle", BlockProperties{
		"pickles": strconv.Itoa(b.Pickles),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b SeaPickle) New(props BlockProperties) Block {
	return SeaPickle{
		Waterlogged: props["waterlogged"] != "false",
		Pickles: atoi(props["pickles"]),
	}
}