package block

type AttachedPumpkinStem struct {
	Facing string
}

func (b AttachedPumpkinStem) Encode() (string, BlockProperties) {
	return "minecraft:attached_pumpkin_stem", BlockProperties{
		"facing": b.Facing,
	}
}

func (b AttachedPumpkinStem) New(props BlockProperties) Block {
	return AttachedPumpkinStem{
		Facing: props["facing"],
	}
}