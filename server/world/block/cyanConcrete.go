package block

type CyanConcrete struct {
}

func (b CyanConcrete) Encode() (string, BlockProperties) {
	return "minecraft:cyan_concrete", BlockProperties{}
}

func (b CyanConcrete) New(props BlockProperties) Block {
	return CyanConcrete{}
}