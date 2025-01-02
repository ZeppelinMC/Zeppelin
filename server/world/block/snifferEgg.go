package block

import (
	"strconv"
)

type SnifferEgg struct {
	Hatch int
}

func (b SnifferEgg) Encode() (string, BlockProperties) {
	return "minecraft:sniffer_egg", BlockProperties{
		"hatch": strconv.Itoa(b.Hatch),
	}
}

func (b SnifferEgg) New(props BlockProperties) Block {
	return SnifferEgg{
		Hatch: atoi(props["hatch"]),
	}
}