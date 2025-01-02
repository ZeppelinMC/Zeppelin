package block

type PurpleConcretePowder struct {
}

func (b PurpleConcretePowder) Encode() (string, BlockProperties) {
	return "minecraft:purple_concrete_powder", BlockProperties{}
}

func (b PurpleConcretePowder) New(props BlockProperties) Block {
	return PurpleConcretePowder{}
}