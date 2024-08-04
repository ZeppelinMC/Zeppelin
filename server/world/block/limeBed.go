package block

import (
	"strconv"
)

type LimeBed struct {
	Facing string
	Occupied bool
	Part string
}

func (b LimeBed) Encode() (string, BlockProperties) {
	return "minecraft:lime_bed", BlockProperties{
		"facing": b.Facing,
		"occupied": strconv.FormatBool(b.Occupied),
		"part": b.Part,
	}
}

func (b LimeBed) New(props BlockProperties) Block {
	return LimeBed{
		Facing: props["facing"],
		Occupied: props["occupied"] != "false",
		Part: props["part"],
	}
}