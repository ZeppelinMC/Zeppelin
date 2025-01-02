package block

import (
	"strconv"
)

type SkeletonSkull struct {
	Rotation int
	Powered bool
}

func (b SkeletonSkull) Encode() (string, BlockProperties) {
	return "minecraft:skeleton_skull", BlockProperties{
		"rotation": strconv.Itoa(b.Rotation),
		"powered": strconv.FormatBool(b.Powered),
	}
}

func (b SkeletonSkull) New(props BlockProperties) Block {
	return SkeletonSkull{
		Powered: props["powered"] != "false",
		Rotation: atoi(props["rotation"]),
	}
}