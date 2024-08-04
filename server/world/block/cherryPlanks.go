package block

type CherryPlanks struct {
}

func (b CherryPlanks) Encode() (string, BlockProperties) {
	return "minecraft:cherry_planks", BlockProperties{}
}

func (b CherryPlanks) New(props BlockProperties) Block {
	return CherryPlanks{}
}