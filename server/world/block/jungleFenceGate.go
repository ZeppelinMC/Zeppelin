package block

import (
	"strconv"
)

type JungleFenceGate struct {
	Facing string
	InWall bool
	Open bool
	Powered bool
}

func (b JungleFenceGate) Encode() (string, BlockProperties) {
	return "minecraft:jungle_fence_gate", BlockProperties{
		"facing": b.Facing,
		"in_wall": strconv.FormatBool(b.InWall),
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b JungleFenceGate) New(props BlockProperties) Block {
	return JungleFenceGate{
		InWall: props["in_wall"] != "false",
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Facing: props["facing"],
	}
}