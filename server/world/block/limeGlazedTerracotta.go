package block

type LimeGlazedTerracotta struct {
	Facing string
}

func (b LimeGlazedTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:lime_glazed_terracotta", BlockProperties{
		"facing": b.Facing,
	}
}

func (b LimeGlazedTerracotta) New(props BlockProperties) Block {
	return LimeGlazedTerracotta{
		Facing: props["facing"],
	}
}