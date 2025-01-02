package block

type SprucePlanks struct {
}

func (b SprucePlanks) Encode() (string, BlockProperties) {
	return "minecraft:spruce_planks", BlockProperties{}
}

func (b SprucePlanks) New(props BlockProperties) Block {
	return SprucePlanks{}
}