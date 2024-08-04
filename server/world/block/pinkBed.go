package block

import (
	"strconv"
)

type PinkBed struct {
	Facing string
	Occupied bool
	Part string
}

func (b PinkBed) Encode() (string, BlockProperties) {
	return "minecraft:pink_bed", BlockProperties{
		"occupied": strconv.FormatBool(b.Occupied),
		"part": b.Part,
		"facing": b.Facing,
	}
}

func (b PinkBed) New(props BlockProperties) Block {
	return PinkBed{
		Part: props["part"],
		Facing: props["facing"],
		Occupied: props["occupied"] != "false",
	}
}