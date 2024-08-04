package block

type OrangeGlazedTerracotta struct {
	Facing string
}

func (b OrangeGlazedTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:orange_glazed_terracotta", BlockProperties{
		"facing": b.Facing,
	}
}

func (b OrangeGlazedTerracotta) New(props BlockProperties) Block {
	return OrangeGlazedTerracotta{
		Facing: props["facing"],
	}
}