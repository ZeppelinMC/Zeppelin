package block

type LimeShulkerBox struct {
	Facing string
}

func (b LimeShulkerBox) Encode() (string, BlockProperties) {
	return "minecraft:lime_shulker_box", BlockProperties{
		"facing": b.Facing,
	}
}

func (b LimeShulkerBox) New(props BlockProperties) Block {
	return LimeShulkerBox{
		Facing: props["facing"],
	}
}