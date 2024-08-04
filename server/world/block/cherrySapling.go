package block

import (
	"strconv"
)

type CherrySapling struct {
	Stage int
}

func (b CherrySapling) Encode() (string, BlockProperties) {
	return "minecraft:cherry_sapling", BlockProperties{
		"stage": strconv.Itoa(b.Stage),
	}
}

func (b CherrySapling) New(props BlockProperties) Block {
	return CherrySapling{
		Stage: atoi(props["stage"]),
	}
}