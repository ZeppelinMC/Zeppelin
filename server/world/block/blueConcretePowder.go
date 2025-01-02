package block

type BlueConcretePowder struct {
}

func (b BlueConcretePowder) Encode() (string, BlockProperties) {
	return "minecraft:blue_concrete_powder", BlockProperties{}
}

func (b BlueConcretePowder) New(props BlockProperties) Block {
	return BlueConcretePowder{}
}