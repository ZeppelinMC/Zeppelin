package block

type YellowGlazedTerracotta struct {
	Facing string
}

func (b YellowGlazedTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:yellow_glazed_terracotta", BlockProperties{
		"facing": b.Facing,
	}
}

func (b YellowGlazedTerracotta) New(props BlockProperties) Block {
	return YellowGlazedTerracotta{
		Facing: props["facing"],
	}
}