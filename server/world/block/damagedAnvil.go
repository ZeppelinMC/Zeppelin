package block

type DamagedAnvil struct {
	Facing string
}

func (b DamagedAnvil) Encode() (string, BlockProperties) {
	return "minecraft:damaged_anvil", BlockProperties{
		"facing": b.Facing,
	}
}

func (b DamagedAnvil) New(props BlockProperties) Block {
	return DamagedAnvil{
		Facing: props["facing"],
	}
}