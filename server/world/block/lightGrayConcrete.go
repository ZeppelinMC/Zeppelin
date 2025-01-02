package block

type LightGrayConcrete struct {
}

func (b LightGrayConcrete) Encode() (string, BlockProperties) {
	return "minecraft:light_gray_concrete", BlockProperties{}
}

func (b LightGrayConcrete) New(props BlockProperties) Block {
	return LightGrayConcrete{}
}