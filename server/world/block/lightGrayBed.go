package block

import (
	"strconv"
)

type LightGrayBed struct {
	Facing string
	Occupied bool
	Part string
}

func (b LightGrayBed) Encode() (string, BlockProperties) {
	return "minecraft:light_gray_bed", BlockProperties{
		"facing": b.Facing,
		"occupied": strconv.FormatBool(b.Occupied),
		"part": b.Part,
	}
}

func (b LightGrayBed) New(props BlockProperties) Block {
	return LightGrayBed{
		Facing: props["facing"],
		Occupied: props["occupied"] != "false",
		Part: props["part"],
	}
}