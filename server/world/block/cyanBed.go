package block

import (
	"strconv"
)

type CyanBed struct {
	Occupied bool
	Part string
	Facing string
}

func (b CyanBed) Encode() (string, BlockProperties) {
	return "minecraft:cyan_bed", BlockProperties{
		"facing": b.Facing,
		"occupied": strconv.FormatBool(b.Occupied),
		"part": b.Part,
	}
}

func (b CyanBed) New(props BlockProperties) Block {
	return CyanBed{
		Facing: props["facing"],
		Occupied: props["occupied"] != "false",
		Part: props["part"],
	}
}