package block

type LightBlueConcretePowder struct {
}

func (b LightBlueConcretePowder) Encode() (string, BlockProperties) {
	return "minecraft:light_blue_concrete_powder", BlockProperties{}
}

func (b LightBlueConcretePowder) New(props BlockProperties) Block {
	return LightBlueConcretePowder{}
}