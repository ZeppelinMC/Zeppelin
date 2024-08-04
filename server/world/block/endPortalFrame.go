package block

import (
	"strconv"
)

type EndPortalFrame struct {
	Eye bool
	Facing string
}

func (b EndPortalFrame) Encode() (string, BlockProperties) {
	return "minecraft:end_portal_frame", BlockProperties{
		"eye": strconv.FormatBool(b.Eye),
		"facing": b.Facing,
	}
}

func (b EndPortalFrame) New(props BlockProperties) Block {
	return EndPortalFrame{
		Eye: props["eye"] != "false",
		Facing: props["facing"],
	}
}