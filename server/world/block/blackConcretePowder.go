package block

type BlackConcretePowder struct {
}

func (b BlackConcretePowder) Encode() (string, BlockProperties) {
	return "minecraft:black_concrete_powder", BlockProperties{}
}

func (b BlackConcretePowder) New(props BlockProperties) Block {
	return BlackConcretePowder{}
}