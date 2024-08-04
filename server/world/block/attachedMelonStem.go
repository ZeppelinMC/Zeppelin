package block

type AttachedMelonStem struct {
	Facing string
}

func (b AttachedMelonStem) Encode() (string, BlockProperties) {
	return "minecraft:attached_melon_stem", BlockProperties{
		"facing": b.Facing,
	}
}

func (b AttachedMelonStem) New(props BlockProperties) Block {
	return AttachedMelonStem{
		Facing: props["facing"],
	}
}