package block

import (
	"strconv"
)

type WitherSkeletonWallSkull struct {
	Facing string
	Powered bool
}

func (b WitherSkeletonWallSkull) Encode() (string, BlockProperties) {
	return "minecraft:wither_skeleton_wall_skull", BlockProperties{
		"facing": b.Facing,
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b WitherSkeletonWallSkull) New(props BlockProperties) Block {
	return WitherSkeletonWallSkull{
		Facing: props["facing"],
		Powered: props["powered"] != "false",
	}
}