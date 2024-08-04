package block

import (
	"strconv"
)

type SpruceFenceGate struct {
	Open bool
	Powered bool
	Facing string
	InWall bool
}

func (b SpruceFenceGate) Encode() (string, BlockProperties) {
	return "minecraft:spruce_fence_gate", BlockProperties{
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
		"in_wall": strconv.FormatBool(b.InWall),
	}
}

func (b SpruceFenceGate) New(props BlockProperties) Block {
	return SpruceFenceGate{
		Facing: props["facing"],
		InWall: props["in_wall"] != "false",
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
	}
}