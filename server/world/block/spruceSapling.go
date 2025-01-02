package block

import (
	"strconv"
)

type SpruceSapling struct {
	Stage int
}

func (b SpruceSapling) Encode() (string, BlockProperties) {
	return "minecraft:spruce_sapling", BlockProperties{
		"stage": strconv.Itoa(b.Stage),
	}
}

func (b SpruceSapling) New(props BlockProperties) Block {
	return SpruceSapling{
		Stage: atoi(props["stage"]),
	}
}