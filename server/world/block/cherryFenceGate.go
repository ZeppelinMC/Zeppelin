package block

import (
	"strconv"
)

type CherryFenceGate struct {
	Facing string
	InWall bool
	Open bool
	Powered bool
}

func (b CherryFenceGate) Encode() (string, BlockProperties) {
	return "minecraft:cherry_fence_gate", BlockProperties{
		"facing": b.Facing,
		"in_wall": strconv.FormatBool(b.InWall),
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b CherryFenceGate) New(props BlockProperties) Block {
	return CherryFenceGate{
		InWall: props["in_wall"] != "false",
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Facing: props["facing"],
	}
}