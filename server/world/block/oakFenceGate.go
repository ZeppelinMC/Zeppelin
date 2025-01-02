package block

import (
	"strconv"
)

type OakFenceGate struct {
	InWall bool
	Open bool
	Powered bool
	Facing string
}

func (b OakFenceGate) Encode() (string, BlockProperties) {
	return "minecraft:oak_fence_gate", BlockProperties{
		"facing": b.Facing,
		"in_wall": strconv.FormatBool(b.InWall),
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b OakFenceGate) New(props BlockProperties) Block {
	return OakFenceGate{
		Facing: props["facing"],
		InWall: props["in_wall"] != "false",
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
	}
}