package block

type Andesite struct {
}

func (b Andesite) Encode() (string, BlockProperties) {
	return "minecraft:andesite", BlockProperties{}
}

func (b Andesite) New(props BlockProperties) Block {
	return Andesite{}
}