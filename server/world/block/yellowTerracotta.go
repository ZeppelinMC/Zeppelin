package block

type YellowTerracotta struct {
}

func (b YellowTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:yellow_terracotta", BlockProperties{}
}

func (b YellowTerracotta) New(props BlockProperties) Block {
	return YellowTerracotta{}
}