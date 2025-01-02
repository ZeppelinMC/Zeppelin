package block

import (
	"strconv"
)

type LightBlueBed struct {
	Facing string
	Occupied bool
	Part string
}

func (b LightBlueBed) Encode() (string, BlockProperties) {
	return "minecraft:light_blue_bed", BlockProperties{
		"occupied": strconv.FormatBool(b.Occupied),
		"part": b.Part,
		"facing": b.Facing,
	}
}

func (b LightBlueBed) New(props BlockProperties) Block {
	return LightBlueBed{
		Part: props["part"],
		Facing: props["facing"],
		Occupied: props["occupied"] != "false",
	}
}