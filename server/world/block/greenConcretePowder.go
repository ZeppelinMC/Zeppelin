package block

type GreenConcretePowder struct {
}

func (b GreenConcretePowder) Encode() (string, BlockProperties) {
	return "minecraft:green_concrete_powder", BlockProperties{}
}

func (b GreenConcretePowder) New(props BlockProperties) Block {
	return GreenConcretePowder{}
}