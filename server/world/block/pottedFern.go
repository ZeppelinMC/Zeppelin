package block

type PottedFern struct {
}

func (b PottedFern) Encode() (string, BlockProperties) {
	return "minecraft:potted_fern", BlockProperties{}
}

func (b PottedFern) New(props BlockProperties) Block {
	return PottedFern{}
}