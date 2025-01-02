package block

import (
	"strconv"
)

type StoneButton struct {
	Face string
	Facing string
	Powered bool
}

func (b StoneButton) Encode() (string, BlockProperties) {
	return "minecraft:stone_button", BlockProperties{
		"face": b.Face,
		"facing": b.Facing,
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b StoneButton) New(props BlockProperties) Block {
	return StoneButton{
		Facing: props["facing"],
		Powered: props["powered"] != "false",
		Face: props["face"],
	}
}