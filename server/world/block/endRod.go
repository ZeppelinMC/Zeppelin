package block

type EndRod struct {
	Facing string
}

func (b EndRod) Encode() (string, BlockProperties) {
	return "minecraft:end_rod", BlockProperties{
		"facing": b.Facing,
	}
}

func (b EndRod) New(props BlockProperties) Block {
	return EndRod{
		Facing: props["facing"],
	}
}