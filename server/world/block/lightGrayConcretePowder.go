package block

type LightGrayConcretePowder struct {
}

func (b LightGrayConcretePowder) Encode() (string, BlockProperties) {
	return "minecraft:light_gray_concrete_powder", BlockProperties{}
}

func (b LightGrayConcretePowder) New(props BlockProperties) Block {
	return LightGrayConcretePowder{}
}