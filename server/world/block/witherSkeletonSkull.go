package block

import (
	"strconv"
)

type WitherSkeletonSkull struct {
	Powered bool
	Rotation int
}

func (b WitherSkeletonSkull) Encode() (string, BlockProperties) {
	return "minecraft:wither_skeleton_skull", BlockProperties{
		"powered": strconv.FormatBool(b.Powered),
		"rotation": strconv.Itoa(b.Rotation),
	}
}

func (b WitherSkeletonSkull) New(props BlockProperties) Block {
	return WitherSkeletonSkull{
		Powered: props["powered"] != "false",
		Rotation: atoi(props["rotation"]),
	}
}