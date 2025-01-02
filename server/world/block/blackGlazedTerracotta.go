package block

type BlackGlazedTerracotta struct {
	Facing string
}

func (b BlackGlazedTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:black_glazed_terracotta", BlockProperties{
		"facing": b.Facing,
	}
}

func (b BlackGlazedTerracotta) New(props BlockProperties) Block {
	return BlackGlazedTerracotta{
		Facing: props["facing"],
	}
}