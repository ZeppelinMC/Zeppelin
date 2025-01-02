package block

type BlueTerracotta struct {
}

func (b BlueTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:blue_terracotta", BlockProperties{}
}

func (b BlueTerracotta) New(props BlockProperties) Block {
	return BlueTerracotta{}
}