package block

import (
	"strconv"
)

type YellowBed struct {
	Facing string
	Occupied bool
	Part string
}

func (b YellowBed) Encode() (string, BlockProperties) {
	return "minecraft:yellow_bed", BlockProperties{
		"occupied": strconv.FormatBool(b.Occupied),
		"part": b.Part,
		"facing": b.Facing,
	}
}

func (b YellowBed) New(props BlockProperties) Block {
	return YellowBed{
		Facing: props["facing"],
		Occupied: props["occupied"] != "false",
		Part: props["part"],
	}
}