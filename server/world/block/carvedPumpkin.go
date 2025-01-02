package block

type CarvedPumpkin struct {
	Facing string
}

func (b CarvedPumpkin) Encode() (string, BlockProperties) {
	return "minecraft:carved_pumpkin", BlockProperties{
		"facing": b.Facing,
	}
}

func (b CarvedPumpkin) New(props BlockProperties) Block {
	return CarvedPumpkin{
		Facing: props["facing"],
	}
}