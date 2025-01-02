package block

import (
	"strconv"
)

type BirchFenceGate struct {
	Facing string
	InWall bool
	Open bool
	Powered bool
}

func (b BirchFenceGate) Encode() (string, BlockProperties) {
	return "minecraft:birch_fence_gate", BlockProperties{
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
		"in_wall": strconv.FormatBool(b.InWall),
	}
}

func (b BirchFenceGate) New(props BlockProperties) Block {
	return BirchFenceGate{
		Facing: props["facing"],
		InWall: props["in_wall"] != "false",
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
	}
}