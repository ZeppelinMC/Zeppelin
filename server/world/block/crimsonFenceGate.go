package block

import (
	"strconv"
)

type CrimsonFenceGate struct {
	InWall bool
	Open bool
	Powered bool
	Facing string
}

func (b CrimsonFenceGate) Encode() (string, BlockProperties) {
	return "minecraft:crimson_fence_gate", BlockProperties{
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
		"in_wall": strconv.FormatBool(b.InWall),
	}
}

func (b CrimsonFenceGate) New(props BlockProperties) Block {
	return CrimsonFenceGate{
		Powered: props["powered"] != "false",
		Facing: props["facing"],
		InWall: props["in_wall"] != "false",
		Open: props["open"] != "false",
	}
}