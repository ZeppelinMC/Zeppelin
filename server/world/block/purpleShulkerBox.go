package block

type PurpleShulkerBox struct {
	Facing string
}

func (b PurpleShulkerBox) Encode() (string, BlockProperties) {
	return "minecraft:purple_shulker_box", BlockProperties{
		"facing": b.Facing,
	}
}

func (b PurpleShulkerBox) New(props BlockProperties) Block {
	return PurpleShulkerBox{
		Facing: props["facing"],
	}
}