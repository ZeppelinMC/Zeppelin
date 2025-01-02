package block

type RedShulkerBox struct {
	Facing string
}

func (b RedShulkerBox) Encode() (string, BlockProperties) {
	return "minecraft:red_shulker_box", BlockProperties{
		"facing": b.Facing,
	}
}

func (b RedShulkerBox) New(props BlockProperties) Block {
	return RedShulkerBox{
		Facing: props["facing"],
	}
}