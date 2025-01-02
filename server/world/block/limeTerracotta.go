package block

type LimeTerracotta struct {
}

func (b LimeTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:lime_terracotta", BlockProperties{}
}

func (b LimeTerracotta) New(props BlockProperties) Block {
	return LimeTerracotta{}
}