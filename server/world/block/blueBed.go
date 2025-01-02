package block

import (
	"strconv"
)

type BlueBed struct {
	Part string
	Facing string
	Occupied bool
}

func (b BlueBed) Encode() (string, BlockProperties) {
	return "minecraft:blue_bed", BlockProperties{
		"part": b.Part,
		"facing": b.Facing,
		"occupied": strconv.FormatBool(b.Occupied),
	}
}

func (b BlueBed) New(props BlockProperties) Block {
	return BlueBed{
		Occupied: props["occupied"] != "false",
		Part: props["part"],
		Facing: props["facing"],
	}
}