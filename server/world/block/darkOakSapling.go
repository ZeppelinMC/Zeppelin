package block

import (
	"strconv"
)

type DarkOakSapling struct {
	Stage int
}

func (b DarkOakSapling) Encode() (string, BlockProperties) {
	return "minecraft:dark_oak_sapling", BlockProperties{
		"stage": strconv.Itoa(b.Stage),
	}
}

func (b DarkOakSapling) New(props BlockProperties) Block {
	return DarkOakSapling{
		Stage: atoi(props["stage"]),
	}
}