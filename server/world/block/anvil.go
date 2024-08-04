package block

type Anvil struct {
	Facing string
}

func (b Anvil) Encode() (string, BlockProperties) {
	return "minecraft:anvil", BlockProperties{
		"facing": b.Facing,
	}
}

func (b Anvil) New(props BlockProperties) Block {
	return Anvil{
		Facing: props["facing"],
	}
}