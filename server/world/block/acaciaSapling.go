package block

import (
	"strconv"
)

type AcaciaSapling struct {
	Stage int
}

func (b AcaciaSapling) Encode() (string, BlockProperties) {
	return "minecraft:acacia_sapling", BlockProperties{
		"stage": strconv.Itoa(b.Stage),
	}
}

func (b AcaciaSapling) New(props BlockProperties) Block {
	return AcaciaSapling{
		Stage: atoi(props["stage"]),
	}
}