package block

import (
	"strconv"
)

type PurpleBed struct {
	Facing string
	Occupied bool
	Part string
}

func (b PurpleBed) Encode() (string, BlockProperties) {
	return "minecraft:purple_bed", BlockProperties{
		"facing": b.Facing,
		"occupied": strconv.FormatBool(b.Occupied),
		"part": b.Part,
	}
}

func (b PurpleBed) New(props BlockProperties) Block {
	return PurpleBed{
		Facing: props["facing"],
		Occupied: props["occupied"] != "false",
		Part: props["part"],
	}
}