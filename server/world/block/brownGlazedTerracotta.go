package block

type BrownGlazedTerracotta struct {
	Facing string
}

func (b BrownGlazedTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:brown_glazed_terracotta", BlockProperties{
		"facing": b.Facing,
	}
}

func (b BrownGlazedTerracotta) New(props BlockProperties) Block {
	return BrownGlazedTerracotta{
		Facing: props["facing"],
	}
}