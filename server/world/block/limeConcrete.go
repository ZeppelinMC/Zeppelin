package block

type LimeConcrete struct {
}

func (b LimeConcrete) Encode() (string, BlockProperties) {
	return "minecraft:lime_concrete", BlockProperties{}
}

func (b LimeConcrete) New(props BlockProperties) Block {
	return LimeConcrete{}
}