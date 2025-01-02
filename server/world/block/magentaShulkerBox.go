package block

type MagentaShulkerBox struct {
	Facing string
}

func (b MagentaShulkerBox) Encode() (string, BlockProperties) {
	return "minecraft:magenta_shulker_box", BlockProperties{
		"facing": b.Facing,
	}
}

func (b MagentaShulkerBox) New(props BlockProperties) Block {
	return MagentaShulkerBox{
		Facing: props["facing"],
	}
}