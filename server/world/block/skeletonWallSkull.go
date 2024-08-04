package block

import (
	"strconv"
)

type SkeletonWallSkull struct {
	Powered bool
	Facing string
}

func (b SkeletonWallSkull) Encode() (string, BlockProperties) {
	return "minecraft:skeleton_wall_skull", BlockProperties{
		"facing": b.Facing,
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b SkeletonWallSkull) New(props BlockProperties) Block {
	return SkeletonWallSkull{
		Facing: props["facing"],
		Powered: props["powered"] != "false",
	}
}