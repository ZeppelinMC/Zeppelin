package block

type RedGlazedTerracotta struct {
	Facing string
}

func (b RedGlazedTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:red_glazed_terracotta", BlockProperties{
		"facing": b.Facing,
	}
}

func (b RedGlazedTerracotta) New(props BlockProperties) Block {
	return RedGlazedTerracotta{
		Facing: props["facing"],
	}
}