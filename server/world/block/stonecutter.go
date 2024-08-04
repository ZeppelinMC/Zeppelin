package block

type Stonecutter struct {
	Facing string
}

func (b Stonecutter) Encode() (string, BlockProperties) {
	return "minecraft:stonecutter", BlockProperties{
		"facing": b.Facing,
	}
}

func (b Stonecutter) New(props BlockProperties) Block {
	return Stonecutter{
		Facing: props["facing"],
	}
}