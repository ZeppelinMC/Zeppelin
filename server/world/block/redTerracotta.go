package block

type RedTerracotta struct {
}

func (b RedTerracotta) Encode() (string, BlockProperties) {
	return "minecraft:red_terracotta", BlockProperties{}
}

func (b RedTerracotta) New(props BlockProperties) Block {
	return RedTerracotta{}
}