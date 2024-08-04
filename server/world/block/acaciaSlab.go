package block

import (
	"strconv"
)

type AcaciaSlab struct {
	Type string
	Waterlogged bool
}

func (b AcaciaSlab) Encode() (string, BlockProperties) {
	return "minecraft:acacia_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b AcaciaSlab) New(props BlockProperties) Block {
	return AcaciaSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}