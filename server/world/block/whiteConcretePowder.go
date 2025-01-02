package block

type WhiteConcretePowder struct {
}

func (b WhiteConcretePowder) Encode() (string, BlockProperties) {
	return "minecraft:white_concrete_powder", BlockProperties{}
}

func (b WhiteConcretePowder) New(props BlockProperties) Block {
	return WhiteConcretePowder{}
}