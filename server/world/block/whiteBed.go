package block

import (
	"strconv"
)

type WhiteBed struct {
	Facing string
	Occupied bool
	Part string
}

func (b WhiteBed) Encode() (string, BlockProperties) {
	return "minecraft:white_bed", BlockProperties{
		"part": b.Part,
		"facing": b.Facing,
		"occupied": strconv.FormatBool(b.Occupied),
	}
}

func (b WhiteBed) New(props BlockProperties) Block {
	return WhiteBed{
		Facing: props["facing"],
		Occupied: props["occupied"] != "false",
		Part: props["part"],
	}
}