package block

type WhiteConcrete struct {
}

func (b WhiteConcrete) Encode() (string, BlockProperties) {
	return "minecraft:white_concrete", BlockProperties{}
}

func (b WhiteConcrete) New(props BlockProperties) Block {
	return WhiteConcrete{}
}