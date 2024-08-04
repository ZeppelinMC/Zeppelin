package block

type GrayConcrete struct {
}

func (b GrayConcrete) Encode() (string, BlockProperties) {
	return "minecraft:gray_concrete", BlockProperties{}
}

func (b GrayConcrete) New(props BlockProperties) Block {
	return GrayConcrete{}
}