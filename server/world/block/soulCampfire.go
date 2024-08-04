package block

import (
	"strconv"
)

type SoulCampfire struct {
	Waterlogged bool
	Facing string
	Lit bool
	SignalFire bool
}

func (b SoulCampfire) Encode() (string, BlockProperties) {
	return "minecraft:soul_campfire", BlockProperties{
		"waterlogged": strconv.FormatBool(b.Waterlogged),
		"facing": b.Facing,
		"lit": strconv.FormatBool(b.Lit),
		"signal_fire": strconv.FormatBool(b.SignalFire),
	}
}

func (b SoulCampfire) New(props BlockProperties) Block {
	return SoulCampfire{
		Lit: props["lit"] != "false",
		SignalFire: props["signal_fire"] != "false",
		Waterlogged: props["waterlogged"] != "false",
		Facing: props["facing"],
	}
}