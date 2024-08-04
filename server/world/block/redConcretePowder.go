package block

type RedConcretePowder struct {
}

func (b RedConcretePowder) Encode() (string, BlockProperties) {
	return "minecraft:red_concrete_powder", BlockProperties{}
}

func (b RedConcretePowder) New(props BlockProperties) Block {
	return RedConcretePowder{}
}