package block

type YellowConcrete struct {
}

func (b YellowConcrete) Encode() (string, BlockProperties) {
	return "minecraft:yellow_concrete", BlockProperties{}
}

func (b YellowConcrete) New(props BlockProperties) Block {
	return YellowConcrete{}
}