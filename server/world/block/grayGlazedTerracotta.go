package block

type GrayGlazedTerracotta struct {
	Facing string
}

func (b GrayGlazedTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:gray_glazed_terracotta", BlockProperties{
		"facing": b.Facing,
	}
}

func (b GrayGlazedTerracotta) New(props BlockProperties) Block {
	return GrayGlazedTerracotta{
		Facing: props["facing"],
	}
}