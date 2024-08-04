package block

type CyanTerracotta struct {
}

func (b CyanTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:cyan_terracotta", BlockProperties{}
}

func (b CyanTerracotta) New(props BlockProperties) Block {
	return CyanTerracotta{}
}