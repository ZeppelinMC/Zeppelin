package block

type ChippedAnvil struct {
	Facing string
}

func (b ChippedAnvil) Encode() (string, BlockProperties) {
	return "minecraft:chipped_anvil", BlockProperties{
		"facing": b.Facing,
	}
}

func (b ChippedAnvil) New(props BlockProperties) Block {
	return ChippedAnvil{
		Facing: props["facing"],
	}
}