package block

type PinkConcretePowder struct {
}

func (b PinkConcretePowder) Encode() (string, BlockProperties) {
	return "minecraft:pink_concrete_powder", BlockProperties{}
}

func (b PinkConcretePowder) New(props BlockProperties) Block {
	return PinkConcretePowder{}
}