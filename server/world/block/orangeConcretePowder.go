package block

type OrangeConcretePowder struct {
}

func (b OrangeConcretePowder) Encode() (string, BlockProperties) {
	return "minecraft:orange_concrete_powder", BlockProperties{}
}

func (b OrangeConcretePowder) New(props BlockProperties) Block {
	return OrangeConcretePowder{}
}