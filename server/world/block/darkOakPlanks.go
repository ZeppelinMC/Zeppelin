package block

type DarkOakPlanks struct {
}

func (b DarkOakPlanks) Encode() (string, BlockProperties) {
	return "minecraft:dark_oak_planks", BlockProperties{}
}

func (b DarkOakPlanks) New(props BlockProperties) Block {
	return DarkOakPlanks{}
}