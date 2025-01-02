package block

import (
	"strconv"
)

type GrayBed struct {
	Facing string
	Occupied bool
	Part string
}

func (b GrayBed) Encode() (string, BlockProperties) {
	return "minecraft:gray_bed", BlockProperties{
		"facing": b.Facing,
		"occupied": strconv.FormatBool(b.Occupied),
		"part": b.Part,
	}
}

func (b GrayBed) New(props BlockProperties) Block {
	return GrayBed{
		Facing: props["facing"],
		Occupied: props["occupied"] != "false",
		Part: props["part"],
	}
}