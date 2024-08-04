package block

type WhiteGlazedTerracotta struct {
	Facing string
}

func (b WhiteGlazedTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:white_glazed_terracotta", BlockProperties{
		"facing": b.Facing,
	}
}

func (b WhiteGlazedTerracotta) New(props BlockProperties) Block {
	return WhiteGlazedTerracotta{
		Facing: props["facing"],
	}
}