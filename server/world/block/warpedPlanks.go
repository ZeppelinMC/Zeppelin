package block

type WarpedPlanks struct {
}

func (b WarpedPlanks) Encode() (string, BlockProperties) {
	return "minecraft:warped_planks", BlockProperties{}
}

func (b WarpedPlanks) New(props BlockProperties) Block {
	return WarpedPlanks{}
}