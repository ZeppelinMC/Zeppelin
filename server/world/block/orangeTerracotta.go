package block

type OrangeTerracotta struct {
}

func (b OrangeTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:orange_terracotta", BlockProperties{}
}

func (b OrangeTerracotta) New(props BlockProperties) Block {
	return OrangeTerracotta{}
}