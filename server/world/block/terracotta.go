package block

type Terracotta struct {
}

func (b Terracotta) Encode() (string, BlockProperties) {
	return "minecraft:terracotta", BlockProperties{}
}

func (b Terracotta) New(props BlockProperties) Block {
	return Terracotta{}
}