package block

type WhiteShulkerBox struct {
	Facing string
}

func (b WhiteShulkerBox) Encode() (string, BlockProperties) {
	return "minecraft:white_shulker_box", BlockProperties{
		"facing": b.Facing,
	}
}

func (b WhiteShulkerBox) New(props BlockProperties) Block {
	return WhiteShulkerBox{
		Facing: props["facing"],
	}
}