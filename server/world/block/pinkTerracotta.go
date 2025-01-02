package block

type PinkTerracotta struct {
}

func (b PinkTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:pink_terracotta", BlockProperties{}
}

func (b PinkTerracotta) New(props BlockProperties) Block {
	return PinkTerracotta{}
}