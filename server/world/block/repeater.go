package block

import (
	"strconv"
)

type Repeater struct {
	Locked bool
	Powered bool
	Delay int
	Facing string
}

func (b Repeater) Encode() (string, BlockProperties) {
	return "minecraft:repeater", BlockProperties{
		"locked": strconv.FormatBool(b.Locked),
		"powered": strconv.FormatBool(b.Powered),
		"delay": strconv.Itoa(b.Delay),
		"facing": b.Facing,
	}
}

func (b Repeater) New(props BlockProperties) Block {
	return Repeater{
		Locked: props["locked"] != "false",
		Powered: props["powered"] != "false",
		Delay: atoi(props["delay"]),
		Facing: props["facing"],
	}
}