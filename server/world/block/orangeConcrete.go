package block

type OrangeConcrete struct {
}

func (b OrangeConcrete) Encode() (string, BlockProperties) {
	return "minecraft:orange_concrete", BlockProperties{}
}

func (b OrangeConcrete) New(props BlockProperties) Block {
	return OrangeConcrete{}
}