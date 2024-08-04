package block

type LightGrayGlazedTerracotta struct {
	Facing string
}

func (b LightGrayGlazedTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:light_gray_glazed_terracotta", BlockProperties{
		"facing": b.Facing,
	}
}

func (b LightGrayGlazedTerracotta) New(props BlockProperties) Block {
	return LightGrayGlazedTerracotta{
		Facing: props["facing"],
	}
}