package block

import (
	"strconv"
)

type StickyPiston struct {
	Extended bool
	Facing string
}

func (b StickyPiston) Encode() (string, BlockProperties) {
	return "minecraft:sticky_piston", BlockProperties{
		"extended": strconv.FormatBool(b.Extended),
		"facing": b.Facing,
	}
}

func (b StickyPiston) New(props BlockProperties) Block {
	return StickyPiston{
		Facing: props["facing"],
		Extended: props["extended"] != "false",
	}
}