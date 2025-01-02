package block

type PurpleTerracotta struct {
}

func (b PurpleTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:purple_terracotta", BlockProperties{}
}

func (b PurpleTerracotta) New(props BlockProperties) Block {
	return PurpleTerracotta{}
}