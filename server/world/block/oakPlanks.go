package block

type OakPlanks struct {
}

func (b OakPlanks) Encode() (string, BlockProperties) {
	return "minecraft:oak_planks", BlockProperties{}
}

func (b OakPlanks) New(props BlockProperties) Block {
	return OakPlanks{}
}