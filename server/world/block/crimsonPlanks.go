package block

type CrimsonPlanks struct {
}

func (b CrimsonPlanks) Encode() (string, BlockProperties) {
	return "minecraft:crimson_planks", BlockProperties{}
}

func (b CrimsonPlanks) New(props BlockProperties) Block {
	return CrimsonPlanks{}
}