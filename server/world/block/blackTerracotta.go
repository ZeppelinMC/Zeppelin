package block

type BlackTerracotta struct {
}

func (b BlackTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:black_terracotta", BlockProperties{}
}

func (b BlackTerracotta) New(props BlockProperties) Block {
	return BlackTerracotta{}
}