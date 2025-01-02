package block

type GrayTerracotta struct {
}

func (b GrayTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:gray_terracotta", BlockProperties{}
}

func (b GrayTerracotta) New(props BlockProperties) Block {
	return GrayTerracotta{}
}