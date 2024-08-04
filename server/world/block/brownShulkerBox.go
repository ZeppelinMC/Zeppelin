package block

type BrownShulkerBox struct {
	Facing string
}

func (b BrownShulkerBox) Encode() (string, BlockProperties) {
	return "minecraft:brown_shulker_box", BlockProperties{
		"facing": b.Facing,
	}
}

func (b BrownShulkerBox) New(props BlockProperties) Block {
	return BrownShulkerBox{
		Facing: props["facing"],
	}
}