package block

import (
	"strconv"
)

type PurpurSlab struct {
	Waterlogged bool
	Type string
}

func (b PurpurSlab) Encode() (string, BlockProperties) {
	return "minecraft:purpur_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b PurpurSlab) New(props BlockProperties) Block {
	return PurpurSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}