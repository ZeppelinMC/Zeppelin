package block

type JunglePlanks struct {
}

func (b JunglePlanks) Encode() (string, BlockProperties) {
	return "minecraft:jungle_planks", BlockProperties{}
}

func (b JunglePlanks) New(props BlockProperties) Block {
	return JunglePlanks{}
}