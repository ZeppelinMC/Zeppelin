package block

type PinkShulkerBox struct {
	Facing string
}

func (b PinkShulkerBox) Encode() (string, BlockProperties) {
	return "minecraft:pink_shulker_box", BlockProperties{
		"facing": b.Facing,
	}
}

func (b PinkShulkerBox) New(props BlockProperties) Block {
	return PinkShulkerBox{
		Facing: props["facing"],
	}
}