package block

type PottedSpruceSapling struct {
}

func (b PottedSpruceSapling) Encode() (string, BlockProperties) {
	return "minecraft:potted_spruce_sapling", BlockProperties{}
}

func (b PottedSpruceSapling) New(props BlockProperties) Block {
	return PottedSpruceSapling{}
}