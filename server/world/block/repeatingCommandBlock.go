package block

import (
	"strconv"
)

type RepeatingCommandBlock struct {
	Facing string
	Conditional bool
}

func (b RepeatingCommandBlock) Encode() (string, BlockProperties) {
	return "minecraft:repeating_command_block", BlockProperties{
		"conditional": strconv.FormatBool(b.Conditional),
		"facing": b.Facing,
	}
}

func (b RepeatingCommandBlock) New(props BlockProperties) Block {
	return RepeatingCommandBlock{
		Conditional: props["conditional"] != "false",
		Facing: props["facing"],
	}
}