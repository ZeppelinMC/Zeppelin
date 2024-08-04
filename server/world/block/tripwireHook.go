package block

import (
	"strconv"
)

type TripwireHook struct {
	Powered bool
	Attached bool
	Facing string
}

func (b TripwireHook) Encode() (string, BlockProperties) {
	return "minecraft:tripwire_hook", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"attached": strconv.FormatBool(b.Attached),
		"facing": b.Facing,
	}
}

func (b TripwireHook) New(props BlockProperties) Block {
	return TripwireHook{
		Facing: props["facing"],
		Powered: props["powered"] != "false",
		Attached: props["attached"] != "false",
	}
}