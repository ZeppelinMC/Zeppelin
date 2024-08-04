package block

type MagentaGlazedTerracotta struct {
	Facing string
}

func (b MagentaGlazedTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:magenta_glazed_terracotta", BlockProperties{
		"facing": b.Facing,
	}
}

func (b MagentaGlazedTerracotta) New(props BlockProperties) Block {
	return MagentaGlazedTerracotta{
		Facing: props["facing"],
	}
}