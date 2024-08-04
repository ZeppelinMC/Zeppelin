package block

type PinkGlazedTerracotta struct {
	Facing string
}

func (b PinkGlazedTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:pink_glazed_terracotta", BlockProperties{
		"facing": b.Facing,
	}
}

func (b PinkGlazedTerracotta) New(props BlockProperties) Block {
	return PinkGlazedTerracotta{
		Facing: props["facing"],
	}
}