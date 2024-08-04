package block

import (
	"strconv"
)

type BrownBed struct {
	Facing string
	Occupied bool
	Part string
}

func (b BrownBed) Encode() (string, BlockProperties) {
	return "minecraft:brown_bed", BlockProperties{
		"facing": b.Facing,
		"occupied": strconv.FormatBool(b.Occupied),
		"part": b.Part,
	}
}

func (b BrownBed) New(props BlockProperties) Block {
	return BrownBed{
		Facing: props["facing"],
		Occupied: props["occupied"] != "false",
		Part: props["part"],
	}
}