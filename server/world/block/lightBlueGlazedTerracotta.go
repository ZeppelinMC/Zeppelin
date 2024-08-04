package block

type LightBlueGlazedTerracotta struct {
	Facing string
}

func (b LightBlueGlazedTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:light_blue_glazed_terracotta", BlockProperties{
		"facing": b.Facing,
	}
}

func (b LightBlueGlazedTerracotta) New(props BlockProperties) Block {
	return LightBlueGlazedTerracotta{
		Facing: props["facing"],
	}
}