package block

type PinkConcrete struct {
}

func (b PinkConcrete) Encode() (string, BlockProperties) {
	return "minecraft:pink_concrete", BlockProperties{}
}

func (b PinkConcrete) New(props BlockProperties) Block {
	return PinkConcrete{}
}