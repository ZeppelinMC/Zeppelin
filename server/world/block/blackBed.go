package block

import (
	"strconv"
)

type BlackBed struct {
	Occupied bool
	Part string
	Facing string
}

func (b BlackBed) Encode() (string, BlockProperties) {
	return "minecraft:black_bed", BlockProperties{
		"facing": b.Facing,
		"occupied": strconv.FormatBool(b.Occupied),
		"part": b.Part,
	}
}

func (b BlackBed) New(props BlockProperties) Block {
	return BlackBed{
		Facing: props["facing"],
		Occupied: props["occupied"] != "false",
		Part: props["part"],
	}
}