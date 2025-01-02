package block

type GrayShulkerBox struct {
	Facing string
}

func (b GrayShulkerBox) Encode() (string, BlockProperties) {
	return "minecraft:gray_shulker_box", BlockProperties{
		"facing": b.Facing,
	}
}

func (b GrayShulkerBox) New(props BlockProperties) Block {
	return GrayShulkerBox{
		Facing: props["facing"],
	}
}