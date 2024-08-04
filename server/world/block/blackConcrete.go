package block

type BlackConcrete struct {
}

func (b BlackConcrete) Encode() (string, BlockProperties) {
	return "minecraft:black_concrete", BlockProperties{}
}

func (b BlackConcrete) New(props BlockProperties) Block {
	return BlackConcrete{}
}