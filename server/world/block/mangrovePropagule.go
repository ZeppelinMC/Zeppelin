package block

import (
	"strconv"
)

type MangrovePropagule struct {
	Waterlogged bool
	Age int
	Hanging bool
	Stage int
}

func (b MangrovePropagule) Encode() (string, BlockProperties) {
	return "minecraft:mangrove_propagule", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"age": strconv.Itoa(b.Age),
		"hanging": strconv.FormatBool(b.Hanging),
		"stage": strconv.Itoa(b.Stage),
	}
}

func (b MangrovePropagule) New(props BlockProperties) Block {
	return MangrovePropagule{
		Age: atoi(props["age"]),
		Hanging: props["hanging"] != "false",
		Stage: atoi(props["stage"]),
		Waterlogged: props["waterlogged"] != "false",
	}
}