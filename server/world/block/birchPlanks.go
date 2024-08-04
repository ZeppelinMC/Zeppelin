package block

type BirchPlanks struct {
}

func (b BirchPlanks) Encode() (string, BlockProperties) {
	return "minecraft:birch_planks", BlockProperties{}
}

func (b BirchPlanks) New(props BlockProperties) Block {
	return BirchPlanks{}
}