package block

import (
	"strconv"
)

type OrangeBed struct {
	Part string
	Facing string
	Occupied bool
}

func (b OrangeBed) Encode() (string, BlockProperties) {
	return "minecraft:orange_bed", BlockProperties{
		"facing": b.Facing,
		"occupied": strconv.FormatBool(b.Occupied),
		"part": b.Part,
	}
}

func (b OrangeBed) New(props BlockProperties) Block {
	return OrangeBed{
		Occupied: props["occupied"] != "false",
		Part: props["part"],
		Facing: props["facing"],
	}
}