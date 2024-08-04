package block

type PurpleGlazedTerracotta struct {
	Facing string
}

func (b PurpleGlazedTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:purple_glazed_terracotta", BlockProperties{
		"facing": b.Facing,
	}
}

func (b PurpleGlazedTerracotta) New(props BlockProperties) Block {
	return PurpleGlazedTerracotta{
		Facing: props["facing"],
	}
}