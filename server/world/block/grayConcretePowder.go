package block

type GrayConcretePowder struct {
}

func (b GrayConcretePowder) Encode() (string, BlockProperties) {
	return "minecraft:gray_concrete_powder", BlockProperties{}
}

func (b GrayConcretePowder) New(props BlockProperties) Block {
	return GrayConcretePowder{}
}