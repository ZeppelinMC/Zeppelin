package block

type PottedPoppy struct {
}

func (b PottedPoppy) Encode() (string, BlockProperties) {
	return "minecraft:potted_poppy", BlockProperties{}
}

func (b PottedPoppy) New(props BlockProperties) Block {
	return PottedPoppy{}
}