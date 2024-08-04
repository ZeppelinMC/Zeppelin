package block

type LimeConcretePowder struct {
}

func (b LimeConcretePowder) Encode() (string, BlockProperties) {
	return "minecraft:lime_concrete_powder", BlockProperties{}
}

func (b LimeConcretePowder) New(props BlockProperties) Block {
	return LimeConcretePowder{}
}