package block

type Prismarine struct {
}

func (b Prismarine) Encode() (string, BlockProperties) {
	return "minecraft:prismarine", BlockProperties{}
}

func (b Prismarine) New(props BlockProperties) Block {
	return Prismarine{}
}