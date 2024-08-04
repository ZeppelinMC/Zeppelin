package block

type PottedCactus struct {
}

func (b PottedCactus) Encode() (string, BlockProperties) {
	return "minecraft:potted_cactus", BlockProperties{}
}

func (b PottedCactus) New(props BlockProperties) Block {
	return PottedCactus{}
}