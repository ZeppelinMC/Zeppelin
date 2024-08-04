package block

type BlackShulkerBox struct {
	Facing string
}

func (b BlackShulkerBox) Encode() (string, BlockProperties) {
	return "minecraft:black_shulker_box", BlockProperties{
		"facing": b.Facing,
	}
}

func (b BlackShulkerBox) New(props BlockProperties) Block {
	return BlackShulkerBox{
		Facing: props["facing"],
	}
}