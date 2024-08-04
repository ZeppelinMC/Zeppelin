package block

type LightBlueShulkerBox struct {
	Facing string
}

func (b LightBlueShulkerBox) Encode() (string, BlockProperties) {
	return "minecraft:light_blue_shulker_box", BlockProperties{
		"facing": b.Facing,
	}
}

func (b LightBlueShulkerBox) New(props BlockProperties) Block {
	return LightBlueShulkerBox{
		Facing: props["facing"],
	}
}