package block

type BrownConcretePowder struct {
}

func (b BrownConcretePowder) Encode() (string, BlockProperties) {
	return "minecraft:brown_concrete_powder", BlockProperties{}
}

func (b BrownConcretePowder) New(props BlockProperties) Block {
	return BrownConcretePowder{}
}