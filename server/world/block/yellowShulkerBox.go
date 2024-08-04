package block

type YellowShulkerBox struct {
	Facing string
}

func (b YellowShulkerBox) Encode() (string, BlockProperties) {
	return "minecraft:yellow_shulker_box", BlockProperties{
		"facing": b.Facing,
	}
}

func (b YellowShulkerBox) New(props BlockProperties) Block {
	return YellowShulkerBox{
		Facing: props["facing"],
	}
}