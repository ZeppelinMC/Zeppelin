package block

type PolishedGranite struct {
}

func (b PolishedGranite) Encode() (string, BlockProperties) {
	return "minecraft:polished_granite", BlockProperties{}
}

func (b PolishedGranite) New(props BlockProperties) Block {
	return PolishedGranite{}
}