package block

type GreenGlazedTerracotta struct {
	Facing string
}

func (b GreenGlazedTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:green_glazed_terracotta", BlockProperties{
		"facing": b.Facing,
	}
}

func (b GreenGlazedTerracotta) New(props BlockProperties) Block {
	return GreenGlazedTerracotta{
		Facing: props["facing"],
	}
}