package block

type WhiteTerracotta struct {
}

func (b WhiteTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:white_terracotta", BlockProperties{}
}

func (b WhiteTerracotta) New(props BlockProperties) Block {
	return WhiteTerracotta{}
}