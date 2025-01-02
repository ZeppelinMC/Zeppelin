package block

type BlueGlazedTerracotta struct {
	Facing string
}

func (b BlueGlazedTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:blue_glazed_terracotta", BlockProperties{
		"facing": b.Facing,
	}
}

func (b BlueGlazedTerracotta) New(props BlockProperties) Block {
	return BlueGlazedTerracotta{
		Facing: props["facing"],
	}
}