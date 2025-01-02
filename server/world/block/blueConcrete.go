package block

type BlueConcrete struct {
}

func (b BlueConcrete) Encode() (string, BlockProperties) {
	return "minecraft:blue_concrete", BlockProperties{}
}

func (b BlueConcrete) New(props BlockProperties) Block {
	return BlueConcrete{}
}