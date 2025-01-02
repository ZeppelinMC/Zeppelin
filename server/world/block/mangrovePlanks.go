package block

type MangrovePlanks struct {
}

func (b MangrovePlanks) Encode() (string, BlockProperties) {
	return "minecraft:mangrove_planks", BlockProperties{}
}

func (b MangrovePlanks) New(props BlockProperties) Block {
	return MangrovePlanks{}
}