package block

type GreenShulkerBox struct {
	Facing string
}

func (b GreenShulkerBox) Encode() (string, BlockProperties) {
	return "minecraft:green_shulker_box", BlockProperties{
		"facing": b.Facing,
	}
}

func (b GreenShulkerBox) New(props BlockProperties) Block {
	return GreenShulkerBox{
		Facing: props["facing"],
	}
}