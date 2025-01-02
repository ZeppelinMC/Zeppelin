package block

type MovingPiston struct {
	Type string
	Facing string
}

func (b MovingPiston) Encode() (string, BlockProperties) {
	return "minecraft:moving_piston", BlockProperties{
		"type": b.Type,
		"facing": b.Facing,
	}
}

func (b MovingPiston) New(props BlockProperties) Block {
	return MovingPiston{
		Facing: props["facing"],
		Type: props["type"],
	}
}