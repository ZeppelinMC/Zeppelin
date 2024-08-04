package block

type LightBlueConcrete struct {
}

func (b LightBlueConcrete) Encode() (string, BlockProperties) {
	return "minecraft:light_blue_concrete", BlockProperties{}
}

func (b LightBlueConcrete) New(props BlockProperties) Block {
	return LightBlueConcrete{}
}