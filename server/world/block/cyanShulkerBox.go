package block

type CyanShulkerBox struct {
	Facing string
}

func (b CyanShulkerBox) Encode() (string, BlockProperties) {
	return "minecraft:cyan_shulker_box", BlockProperties{
		"facing": b.Facing,
	}
}

func (b CyanShulkerBox) New(props BlockProperties) Block {
	return CyanShulkerBox{
		Facing: props["facing"],
	}
}