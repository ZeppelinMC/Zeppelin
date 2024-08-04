package block

import (
	"strconv"
)

type ChainCommandBlock struct {
	Facing string
	Conditional bool
}

func (b ChainCommandBlock) Encode() (string, BlockProperties) {
	return "minecraft:chain_command_block", BlockProperties{
		"conditional": strconv.FormatBool(b.Conditional),
		"facing": b.Facing,
	}
}

func (b ChainCommandBlock) New(props BlockProperties) Block {
	return ChainCommandBlock{
		Conditional: props["conditional"] != "false",
		Facing: props["facing"],
	}
}