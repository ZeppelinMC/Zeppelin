package block

import (
	"strconv"
)

type PitcherCrop struct {
	Half string
	Age int
}

func (b PitcherCrop) Encode() (string, BlockProperties) {
	return "minecraft:pitcher_crop", BlockProperties{
		"half": b.Half,
		"age": strconv.Itoa(b.Age),
	}
}

func (b PitcherCrop) New(props BlockProperties) Block {
	return PitcherCrop{
		Half: props["half"],
		Age: atoi(props["age"]),
	}
}