package block

type CyanConcretePowder struct {
}

func (b CyanConcretePowder) Encode() (string, BlockProperties) {
	return "minecraft:cyan_concrete_powder", BlockProperties{}
}

func (b CyanConcretePowder) New(props BlockProperties) Block {
	return CyanConcretePowder{}
}