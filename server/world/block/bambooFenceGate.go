package block

import (
	"strconv"
)

type BambooFenceGate struct {
	Open bool
	Powered bool
	Facing string
	InWall bool
}

func (b BambooFenceGate) Encode() (string, BlockProperties) {
	return "minecraft:bamboo_fence_gate", BlockProperties{
		"facing": b.Facing,
		"in_wall": strconv.FormatBool(b.InWall),
		"open": strconv.FormatBool(b.Open),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b BambooFenceGate) New(props BlockProperties) Block {
	return BambooFenceGate{
		Open: props["open"] != "false",
		Powered: props["powered"] != "false",
		Facing: props["facing"],
		InWall: props["in_wall"] != "false",
	}
}