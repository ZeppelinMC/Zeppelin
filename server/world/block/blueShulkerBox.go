package block

type BlueShulkerBox struct {
	Facing string
}

func (b BlueShulkerBox) Encode() (string, BlockProperties) {
	return "minecraft:blue_shulker_box", BlockProperties{
		"facing": b.Facing,
	}
}

func (b BlueShulkerBox) New(props BlockProperties) Block {
	return BlueShulkerBox{
		Facing: props["facing"],
	}
}