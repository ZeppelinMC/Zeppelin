package block

type YellowConcretePowder struct {
}

func (b YellowConcretePowder) Encode() (string, BlockProperties) {
	return "minecraft:yellow_concrete_powder", BlockProperties{}
}

func (b YellowConcretePowder) New(props BlockProperties) Block {
	return YellowConcretePowder{}
}