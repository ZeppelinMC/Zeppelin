package block

type LightGrayShulkerBox struct {
	Facing string
}

func (b LightGrayShulkerBox) Encode() (string, BlockProperties) {
	return "minecraft:light_gray_shulker_box", BlockProperties{
		"facing": b.Facing,
	}
}

func (b LightGrayShulkerBox) New(props BlockProperties) Block {
	return LightGrayShulkerBox{
		Facing: props["facing"],
	}
}