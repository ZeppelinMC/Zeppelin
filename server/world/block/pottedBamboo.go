package block

type PottedBamboo struct {
}

func (b PottedBamboo) Encode() (string, BlockProperties) {
	return "minecraft:potted_bamboo", BlockProperties{}
}

func (b PottedBamboo) New(props BlockProperties) Block {
	return PottedBamboo{}
}