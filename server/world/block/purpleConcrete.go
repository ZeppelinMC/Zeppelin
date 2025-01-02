package block

type PurpleConcrete struct {
}

func (b PurpleConcrete) Encode() (string, BlockProperties) {
	return "minecraft:purple_concrete", BlockProperties{}
}

func (b PurpleConcrete) New(props BlockProperties) Block {
	return PurpleConcrete{}
}