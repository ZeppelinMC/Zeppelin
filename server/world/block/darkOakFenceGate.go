package block

import (
	"strconv"
)

type DarkOakFenceGate struct {
	Powered bool
	Facing string
	InWall bool
	Open bool
}

func (b DarkOakFenceGate) Encode() (string, BlockProperties) {
	return "minecraft:dark_oak_fence_gate", BlockProperties{
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
		"in_wall": strconv.FormatBool(b.InWall),
	}
}

func (b DarkOakFenceGate) New(props BlockProperties) Block {
	return DarkOakFenceGate{
		Powered: props["powered"] != "false",
		Facing: props["facing"],
		InWall: props["in_wall"] != "false",
		Open: props["open"] != "false",
	}
}