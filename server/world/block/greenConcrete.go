package block

type GreenConcrete struct {
}

func (b GreenConcrete) Encode() (string, BlockProperties) {
	return "minecraft:green_concrete", BlockProperties{}
}

func (b GreenConcrete) New(props BlockProperties) Block {
	return GreenConcrete{}
}