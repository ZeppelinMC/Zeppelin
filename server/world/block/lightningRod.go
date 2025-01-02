package block

import (
	"strconv"
)

type LightningRod struct {
	Facing string
	Powered bool
	Waterlogged bool
}

func (b LightningRod) Encode() (string, BlockProperties) {
	return "minecraft:lightning_rod", BlockProperties{
		"facing": b.Facing,
		"powered": strconv.FormatBool(b.Powered),
		"waterlogged": strconv.FormatBool(b.Waterlogged),
	}
}

func (b LightningRod) New(props BlockProperties) Block {
	return LightningRod{
		Facing: props["facing"],
		Powered: props["powered"] != "false",
		Waterlogged: props["waterlogged"] != "false",
	}
}