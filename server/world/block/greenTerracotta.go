package block

type GreenTerracotta struct {
}

func (b GreenTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:green_terracotta", BlockProperties{}
}

func (b GreenTerracotta) New(props BlockProperties) Block {
	return GreenTerracotta{}
}