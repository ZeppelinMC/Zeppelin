package block

type CyanGlazedTerracotta struct {
	Facing string
}

func (b CyanGlazedTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:cyan_glazed_terracotta", BlockProperties{
		"facing": b.Facing,
	}
}

func (b CyanGlazedTerracotta) New(props BlockProperties) Block {
	return CyanGlazedTerracotta{
		Facing: props["facing"],
	}
}