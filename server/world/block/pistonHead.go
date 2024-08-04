package block

import (
	"strconv"
)

type PistonHead struct {
	Facing string
	Short bool
	Type string
}

func (b PistonHead) Encode() (string, BlockProperties) {
	return "minecraft:piston_head", BlockProperties{
		"facing": b.Facing,
		"short": strconv.FormatBool(b.Short),
		"type": b.Type,
	}
}

func (b PistonHead) New(props BlockProperties) Block {
	return PistonHead{
		Type: props["type"],
		Facing: props["facing"],
		Short: props["short"] != "false",
	}
}