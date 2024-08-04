package block

type RedConcrete struct {
}

func (b RedConcrete) Encode() (string, BlockProperties) {
	return "minecraft:red_concrete", BlockProperties{}
}

func (b RedConcrete) New(props BlockProperties) Block {
	return RedConcrete{}
}