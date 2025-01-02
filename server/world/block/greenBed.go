package block

import (
	"strconv"
)

type GreenBed struct {
	Facing string
	Occupied bool
	Part string
}

func (b GreenBed) Encode() (string, BlockProperties) {
	return "minecraft:green_bed", BlockProperties{
		"part": b.Part,
		"facing": b.Facing,
		"occupied": strconv.FormatBool(b.Occupied),
	}
}

func (b GreenBed) New(props BlockProperties) Block {
	return GreenBed{
		Facing: props["facing"],
		Occupied: props["occupied"] != "false",
		Part: props["part"],
	}
}