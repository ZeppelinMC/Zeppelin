package block

import (
	"strconv"
)

type OakSapling struct {
	Stage int
}

func (b OakSapling) Encode() (string, BlockProperties) {
	return "minecraft:oak_sapling", BlockProperties{
		"stage": strconv.Itoa(b.Stage),
	}
}

func (b OakSapling) New(props BlockProperties) Block {
	return OakSapling{
		Stage: atoi(props["stage"]),
	}
}