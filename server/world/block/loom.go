package block

type Loom struct {
	Facing string
}

func (b Loom) Encode() (string, BlockProperties) {
	return "minecraft:loom", BlockProperties{
		"facing": b.Facing,
	}
}

func (b Loom) New(props BlockProperties) Block {
	return Loom{
		Facing: props["facing"],
	}
}