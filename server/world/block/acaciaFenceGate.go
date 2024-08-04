package block

import (
	"strconv"
)

type AcaciaFenceGate struct {
	InWall bool
	Open bool
	Powered bool
	Facing string
}

func (b AcaciaFenceGate) Encode() (string, BlockProperties) {
	return "minecraft:acacia_fence_gate", BlockProperties{
		"in_wall": strconv.FormatBool(b.InWall),
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
		"facing": b.Facing,
	}
}

func (b AcaciaFenceGate) New(props BlockProperties) Block {
	return AcaciaFenceGate{
		InWall: props["in_wall"] != "false",
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Facing: props["facing"],
	}
}