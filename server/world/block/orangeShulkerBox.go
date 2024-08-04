package block

type OrangeShulkerBox struct {
	Facing string
}

func (b OrangeShulkerBox) Encode() (string, BlockProperties) {
	return "minecraft:orange_shulker_box", BlockProperties{
		"facing": b.Facing,
	}
}

func (b OrangeShulkerBox) New(props BlockProperties) Block {
	return OrangeShulkerBox{
		Facing: props["facing"],
	}
}