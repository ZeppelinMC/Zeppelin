package block

import (
	"strconv"
)

type MangroveFenceGate struct {
	Facing string
	InWall bool
	Open bool
	Powered bool
}

func (b MangroveFenceGate) Encode() (string, BlockProperties) {
	return "minecraft:mangrove_fence_gate", BlockProperties{
		"facing": b.Facing,
		"in_wall": strconv.FormatBool(b.InWall),
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b MangroveFenceGate) New(props BlockProperties) Block {
	return MangroveFenceGate{
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Facing: props["facing"],
		InWall: props["in_wall"] != "false",
	}
}