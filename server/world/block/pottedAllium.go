package block

type PottedAllium struct {
}

func (b PottedAllium) Encode() (string, BlockProperties) {
	return "minecraft:potted_allium", BlockProperties{}
}

func (b PottedAllium) New(props BlockProperties) Block {
	return PottedAllium{}
}