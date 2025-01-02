package block

import (
	"strconv"
)

type BirchSapling struct {
	Stage int
}

func (b BirchSapling) Encode() (string, BlockProperties) {
	return "minecraft:birch_sapling", BlockProperties{
		"stage": strconv.Itoa(b.Stage),
	}
}

func (b BirchSapling) New(props BlockProperties) Block {
	return BirchSapling{
		Stage: atoi(props["stage"]),
	}
}