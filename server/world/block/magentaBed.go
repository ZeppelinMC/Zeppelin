package block

import (
	"strconv"
)

type MagentaBed struct {
	Facing string
	Occupied bool
	Part string
}

func (b MagentaBed) Encode() (string, BlockProperties) {
	return "minecraft:magenta_bed", BlockProperties{
		"occupied": strconv.FormatBool(b.Occupied),
		"part": b.Part,
		"facing": b.Facing,
	}
}

func (b MagentaBed) New(props BlockProperties) Block {
	return MagentaBed{
		Facing: props["facing"],
		Occupied: props["occupied"] != "false",
		Part: props["part"],
	}
}