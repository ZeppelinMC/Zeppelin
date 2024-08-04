package block

import (
	"strconv"
)

type RedBed struct {
	Part string
	Facing string
	Occupied bool
}

func (b RedBed) Encode() (string, BlockProperties) {
	return "minecraft:red_bed", BlockProperties{
		"part": b.Part,
		"facing": b.Facing,
		"occupied": strconv.FormatBool(b.Occupied),
	}
}

func (b RedBed) New(props BlockProperties) Block {
	return RedBed{
		Part: props["part"],
		Facing: props["facing"],
		Occupied: props["occupied"] != "false",
	}
}