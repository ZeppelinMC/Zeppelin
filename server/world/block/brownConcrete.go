package block

type BrownConcrete struct {
}

func (b BrownConcrete) Encode() (string, BlockProperties) {
	return "minecraft:brown_concrete", BlockProperties{}
}

func (b BrownConcrete) New(props BlockProperties) Block {
	return BrownConcrete{}
}