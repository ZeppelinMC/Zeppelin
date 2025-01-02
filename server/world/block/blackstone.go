package block

type Blackstone struct {
}

func (b Blackstone) Encode() (string, BlockProperties) {
	return "minecraft:blackstone", BlockProperties{}
}

func (b Blackstone) New(props BlockProperties) Block {
	return Blackstone{}
}