package block

import (
	"strconv"
)

type MudBrickSlab struct {
	Type string
	Waterlogged bool
}

func (b MudBrickSlab) Encode() (string, BlockProperties) {
	return "minecraft:mud_brick_slab", BlockProperties{
		"type": b.Type,
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b MudBrickSlab) New(props BlockProperties) Block {
	return MudBrickSlab{
		Type: props["type"],
		Waterlogged: props["waterlogged"] != "false",
	}
}