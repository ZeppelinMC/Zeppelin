package block

type BoneBlock struct {
	Axis string
}

func (b BoneBlock) Encode() (string, BlockProperties) {
	return "minecraft:bone_block", BlockProperties{
		"axis": b.Axis,
	}
}

func (b BoneBlock) New(props BlockProperties) Block {
	return BoneBlock{
		Axis: props["axis"],
	}
}