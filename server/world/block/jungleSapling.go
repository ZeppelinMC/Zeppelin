package block

import (
	"strconv"
)

type JungleSapling struct {
	Stage int
}

func (b JungleSapling) Encode() (string, BlockProperties) {
	return "minecraft:jungle_sapling", BlockProperties{
		"stage": strconv.Itoa(b.Stage),
	}
}

func (b JungleSapling) New(props BlockProperties) Block {
	return JungleSapling{
		Stage: atoi(props["stage"]),
	}
}