package block

import (
	"strconv"
)

type WarpedFenceGate struct {
	Facing string
	InWall bool
	Open bool
	Powered bool
}

func (b WarpedFenceGate) Encode() (string, BlockProperties) {
	return "minecraft:warped_fence_gate", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
		"in_wall": strconv.FormatBool(b.InWall),
		"open": strconv.FormatBool(b.Open),
	}
}

func (b WarpedFenceGate) New(props BlockProperties) Block {
	return WarpedFenceGate{
		Facing: props["facing"],
		InWall: props["in_wall"] != "false",
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
	}
}